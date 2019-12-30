package balance

import (
	"api/models"
	"context"
)

type Fetcher interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	ListUserBalances(ctx context.Context) ([]models.UserBalance, error)
	ListTransactions(ctx context.Context) ([]models.Transaction, error)
}

type UseCase interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	ListUserBalances(ctx context.Context) ([]models.UserBalance, error)
	ListTransactions(ctx context.Context) ([]models.Transaction, error)
}
