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

func (s *Service) CheckUserIn(ctx context.Context, user *model.UserWithBalance, at time.Time) (*model.Checkin, error) {
	dateStr := toUTCDateStr(at)
	_, err := model.FindCheckin(ctx, s.db, user.ID, dateStr)
	if err == nil {
		return nil, ErrAlreadyCheckedIn
	} else if err != sql.ErrNoRows {
		return nil, ErrDatabase
	}

	onTime := isOnTime(at)
	ci := &model.Checkin{
		Time:   at,
		UserID: user.ID,
		Date:   dateStr,
		OnTime: onTime,
	}
	if ci.OnTime {
		ci.StarEarned = 1
	}
	if err = ci.Insert(boil.WithDebug(ctx, true), s.db, boil.Infer()); err != nil {
		return nil, ErrDatabase
	}
	return ci, nil
}

func (s *Service) CheckUserOut(ctx context.Context, user *model.UserWithBalance, at time.Time) (*model.Checkin, error) {
	dateStr := toUTCDateStr(at)
	ci, err := model.FindCheckin(ctx, s.db, user.ID, dateStr)
	if err == sql.ErrNoRows {
		return nil, ErrNotCheckedIn
	} else if err != nil {
		return nil, ErrDatabase
	}

	if _, err = ci.Delete(ctx, s.db); err != nil {
		return nil, ErrDatabase
	}
	return ci, nil
}

func (s *Service) ListUserHasTelegramID(ctx context.Context) (model.UserSlice, error) {
	return model.Users(model.UserWhere.TelegramID.IsNotNull()).All(ctx, s.db)
}

func isOnTime(at time.Time) bool {
	y, m, d := at.In(LocalZone).Date()
	dl := time.Date(y, m, d, 9, 31, 0, 0, LocalZone)
	return at.Before(dl)
}

func randomToken() string {
	return xid.New().String()
}

func toUTCDateStr(t time.Time) string {
	return t.In(LocalZone).Format("2006/01/02")
}
