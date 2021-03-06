package telegram

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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

func (s *Service) ListCheckins(ctx context.Context, today time.Time) (model.CheckinSlice, error) {
	dateStr := toUTCDateStr(today)
	cis, err := model.Checkins(
		model.CheckinWhere.Date.EQ(dateStr),
		qm.OrderBy(model.CheckinColumns.Time+" ASC"),
		qm.Load(model.CheckinRels.User),
	).All(ctx, s.db)
	if err != nil {
		log.Err(err).Msg("list checkins failed")
		return nil, ErrDatabase
	}
	return cis, nil
}

func (s *Service) ListStarsInCurrentMonth(ctx context.Context) (model.UserStars, error) {
	qr := `SELECT u.id, u.name, t.stars
FROM users u
LEFT JOIN
	(SELECT ck.user_id as uid, SUM(ck.star_earned) as stars
	FROM checkins ck
	WHERE ck."date" LIKE '` + time.Now().Format("2006/01") + `/%'
	GROUP BY ck.user_id) t ON u.id = t.uid
WHERE u.deleted_at IS NULL AND t.stars IS NOT NULL
ORDER BY t.stars DESC, u.name ASC`
	rows, err := s.db.QueryContext(ctx, qr)
	if err != nil {
		return nil, err
	}

	res := make(model.UserStars, 0)
	for rows.Next() {
		var id string
		var name string
		var stars float64
		if err := rows.Scan(&id, &name, &stars); err != nil {
			continue
		}
		res = append(res, model.UserWithStar{ID: id, Name: name, Stars: stars})
	}
	return res, nil
}

func (s *Service) ListUserNotCheckedInToday(ctx context.Context) ([]model.User, error) {
	cis, err := s.ListCheckins(ctx, time.Now())
	if err != nil {
		return nil, err
	}
	users, err := s.ListUserHasTelegramID(ctx)
	if err != nil {
		return nil, err
	}

	checked := map[string]bool{}
	for _, c := range cis {
		checked[c.UserID] = true
	}

	var us []model.User
	for _, u := range users {
		if !checked[u.ID] && u.RemindCheckin.Bool {
			us = append(us, *u)
		}
	}
	return us, nil
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
