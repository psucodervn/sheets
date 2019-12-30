package repository

import (
	"api/balance"
	"api/config"
	"api/models"
	"context"
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
)

const (
	RangeBalance     = "201911!T1:AF2"
	RangeTransaction = "201911!A5:R"
	RangeUser        = "201911!E2:Q2"
)

type ApiFetcher struct {
	svc     *sheets.SpreadsheetsService
	sheetID string
}

func (f *ApiFetcher) ListUsers(ctx context.Context) ([]models.User, error) {
	userBalances, err := f.ListUserBalances(ctx)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for i := range userBalances {
		users = append(users, userBalances[i].User)
	}
	return users, nil
}

func (f *ApiFetcher) ListTransactions(ctx context.Context) ([]models.Transaction, error) {
	users, err := f.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	rows, err := f.svc.Values.
		Get(f.sheetID, RangeTransaction).
		ValueRenderOption("UNFORMATTED_VALUE").
		Do()
	if err != nil {
		return nil, fmt.Errorf("get value: %w", err)
	}

	log.Debug().Int("transaction_count", len(rows.Values)).Msg("")
	var res []models.Transaction
	for i := range rows.Values {
		raw := rows.Values[i]
		tx, err := toTransaction(raw, users)
		if errors.Is(err, balance.ErrEmptyTransaction) {
			continue
		}
		if err != nil {
			log.Warn().Err(err).Interface("raw", raw).Msg("toTransaction failed")
			continue
		}
		res = append(res, *tx)
	}
	return res, nil
}

func toTransaction(raw []interface{}, users []models.User) (tx *models.Transaction, err error) {
	defer func() {
		if err != nil {
			return
		}
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	totalCount := toFloat64(raw[17])
	totalValue := toFloat64(raw[1])
	if int(totalCount) <= 0 {
		return nil, balance.ErrEmptyTransaction
	}

	// calc senders
	senders := []models.UserTransaction{
		{Name: toString(raw[2]), Amount: toFloat64(raw[1])},
	}
	// calc receivers
	userCount := len(users)
	var receivers []models.UserTransaction
	for i := 0; i < userCount; i++ {
		cnt := toFloat64(raw[i+4])
		if int(cnt) <= 0 {
			continue
		}
		receivers = append(receivers, models.UserTransaction{
			Name:   users[i].Name,
			Amount: totalValue * cnt / totalCount,
		})
	}

	tx = &models.Transaction{
		Description: toString(raw[3]),
		Value:       totalValue,
		Senders:     senders,
		Receivers:   receivers,
	}
	return tx, nil
}

func toString(i interface{}) string {
	if s, ok := i.(string); ok {
		return s
	}
	return ""
}

func toFloat64(i interface{}) float64 {
	if f, ok := i.(float64); ok {
		return f
	}
	return 0
}

func (f *ApiFetcher) ListUserBalances(ctx context.Context) ([]models.UserBalance, error) {
	rows, err := f.svc.Values.
		Get(f.sheetID, RangeBalance).
		ValueRenderOption("UNFORMATTED_VALUE").
		Do()
	if err != nil {
		return nil, fmt.Errorf("get value: %w", err)
	}

	var res []models.UserBalance
	if len(rows.Values) != 2 {
		return nil, balance.ErrInvalidDataFormat
	}

	userCount := len(rows.Values[0])
	for i := 0; i < userCount; i++ {
		name := rows.Values[1][i].(string)
		val := rows.Values[0][i].(float64)
		res = append(res, models.UserBalance{
			User:    models.User{Name: name},
			Balance: models.Balance{Value: val},
		})
	}

	return res, nil
}

func NewApiFetcher(credentialsJSON []byte, sheetID string) *ApiFetcher {
	ctx := context.Background()
	svc, err := sheets.NewService(ctx, option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		log.Fatal().Err(err).Msg("new service failed")
	}
	spreadsheetsService := sheets.NewSpreadsheetsService(svc)
	return &ApiFetcher{svc: spreadsheetsService, sheetID: sheetID}
}

func NewApiFetcherFromEnv() *ApiFetcher {
	var cfg config.GoogleConfig
	envconfig.MustProcess("GOOGLE", &cfg)

	b, err := ioutil.ReadFile(cfg.CredentialsFile)
	if err != nil {
		log.Fatal().Err(err).Msg("read file failed")
	}

	return NewApiFetcher(b, cfg.SheetID)
}
