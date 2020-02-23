package balance

import (
	"api/model"
	"context"
)

type Fetcher interface {
	ListUsers(ctx context.Context) ([]model.User, error)
	ListUserBalances(ctx context.Context) ([]model.UserBalance, error)
	ListTransactions(ctx context.Context) ([]model.Transaction, error)
}

type Service interface {
	ListUsers(ctx context.Context) ([]model.User, error)
	ListUserBalances(ctx context.Context) ([]model.UserBalance, error)
	ListTransactions(ctx context.Context) ([]model.Transaction, error)
}
