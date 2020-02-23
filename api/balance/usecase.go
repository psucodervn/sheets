package balance

import (
	"api/model"
	"context"
)

type BaseService struct {
	fetcher Fetcher
}

func (u *BaseService) ListUsers(ctx context.Context) ([]model.User, error) {
	return u.fetcher.ListUsers(ctx)
}

func (u *BaseService) ListUserBalances(ctx context.Context) ([]model.UserBalance, error) {
	return u.fetcher.ListUserBalances(ctx)
}

func (u *BaseService) ListTransactions(ctx context.Context) ([]model.Transaction, error) {
	return u.fetcher.ListTransactions(ctx)
}

func NewBaseService(fetcher Fetcher) *BaseService {
	return &BaseService{fetcher: fetcher}
}
