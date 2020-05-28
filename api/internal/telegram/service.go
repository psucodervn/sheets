package telegram

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"api/model"
)

type Service struct {
	db      *sql.DB
	botName string
}

func NewService(db *sql.DB, botName string) *Service {
	return &Service{db: db, botName: botName}
}

func (s *Service) GenerateLink(ctx context.Context, userID string) (string, error) {
	token, err := model.TelegramTokens(model.TelegramTokenWhere.UserID.EQ(userID)).One(ctx, s.db)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}
		// token not exists
		token = &model.TelegramToken{
			Token:     randomToken(),
			CreatedAt: time.Now(),
			ExpireAt:  time.Now().Add(time.Hour),
			UserID:    userID,
		}
	} else {
		token.ExpireAt = time.Now().Add(time.Hour)
	}

	if err := token.Upsert(ctx, s.db, true, nil, boil.Infer(), boil.Infer()); err != nil {
		return "", err
	}
	return fmt.Sprintf("https://t.me/%s?start=%s", s.botName, token.Token), nil
}

func (s *Service) AuthTelegramUser(ctx context.Context, token string, recipient string) (*model.User, error) {
	tk, err := model.FindTelegramToken(ctx, s.db, token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTokenInvalid
		}
		return nil, err
	}
	if tk.ExpireAt.Before(time.Now()) {
		_, _ = tk.Delete(ctx, s.db)
		return nil, ErrTokenExpired
	}

	u, err := model.FindUser(ctx, s.db, tk.UserID)
	if err != nil {
		return nil, err
	}

	u.TelegramID = null.StringFrom(recipient)
	if _, err = u.Update(ctx, s.db, boil.Whitelist(model.UserColumns.TelegramID)); err != nil {
		return nil, err
	}
	if _, err = tk.Delete(ctx, s.db); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) GetUserWithBalanceByTelegramID(ctx context.Context, telegramID string) (*model.UserWithBalance, error) {
	u, err := model.UsersWithBalance(model.UserWhere.TelegramID.EQ(null.StringFrom(telegramID))).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func randomToken() string {
	return xid.New().String()
}
