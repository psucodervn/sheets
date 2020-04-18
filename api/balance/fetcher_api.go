package balance

import (
	"api/config"
	"api/model"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	RangeBalance     = "Current!U1:AH2"
	RangeTransaction = "Current!A4:S"
)

type ApiFetcher struct {
	svc     *sheets.SpreadsheetsService
	sheetID string
}

func (f *ApiFetcher) ListUsers(ctx context.Context) ([]model.User, error) {
	rows, err := f.svc.Values.
		Get(f.sheetID, RangeBalance).
		ValueRenderOption("UNFORMATTED_VALUE").
		Do()
	if err != nil {
		return nil, fmt.Errorf("get value: %w", err)
	}

	var res []model.User
	if len(rows.Values) != 2 {
		return nil, ErrInvalidDataFormat
	}

	userCount := len(rows.Values[0])
	for i := 0; i < userCount; i++ {
		name := rows.Values[1][i].(string)
		val := rows.Values[0][i].(float64)
		res = append(res, model.User{
			Name:    name,
			Balance: model.Balance{Val: val},
		})
	}

	return res, nil
}

func (f *ApiFetcher) ListTransactions(ctx context.Context) ([]model.Transaction, error) {
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
	var res []model.Transaction
	t := float64(43894) // magic
	for i := range rows.Values {
		raw := rows.Values[i]
		if toFloat64(raw[0]) > 0 {
			t = toFloat64(raw[0])
		} else {
			raw[0] = t
		}
		tx, err := toTransaction(raw, users)
		if errors.Is(err, ErrEmptyTransaction) {
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

func toTransaction(raw []interface{}, users []model.User) (tx *model.Transaction, err error) {
	defer func() {
		if err != nil {
			return
		}
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	totalCount := toFloat64(raw[18])
	totalValue := toFloat64(raw[1])
	if int(totalCount) <= 0 {
		return nil, ErrEmptyTransaction
	}

	// calc senders
	senders := []model.UserTransaction{
		{Name: toString(raw[2]), Val: toFloat64(raw[1])},
	}
	// calc receivers
	userCount := len(users)
	var receivers []model.UserTransaction
	for i := 0; i < userCount; i++ {
		cnt := toFloat64(raw[i+4])
		if cnt <= 0 {
			continue
		}
		receivers = append(receivers, model.UserTransaction{
			Name: users[i].Name,
			Val:  totalValue * cnt / totalCount,
		})
	}

	tx = &model.Transaction{
		Description: toString(raw[3]),
		TotalValue:  totalValue,
		Senders:     senders,
		Receivers:   receivers,
	}

	begin := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	tx.Time = begin.Add(time.Duration(raw[0].(float64)-2) * 24 * time.Hour)

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

func (f *ApiFetcher) ListUserBalances(ctx context.Context) ([]model.UserBalance, error) {
	rows, err := f.svc.Values.
		Get(f.sheetID, RangeBalance).
		ValueRenderOption("UNFORMATTED_VALUE").
		Do()
	if err != nil {
		return nil, fmt.Errorf("get value: %w", err)
	}

	var res []model.UserBalance
	if len(rows.Values) != 2 {
		return nil, ErrInvalidDataFormat
	}

	userCount := len(rows.Values[0])
	for i := 0; i < userCount; i++ {
		name := rows.Values[1][i].(string)
		val := rows.Values[0][i].(float64)
		res = append(res, model.UserBalance{
			User:    model.User{Name: name},
			Balance: model.Balance{Val: val},
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
	var cfg config.GoogleDocsConfig
	envconfig.MustProcess("GOOGLE", &cfg)

	b, err := ioutil.ReadFile(cfg.CredentialsFile)
	if err != nil {
		log.Fatal().Err(err).Msg("read file failed")
	}

	return NewApiFetcher(b, cfg.SheetID)
}
