package balance

import (
	"api/api"
	"api/model"
	"context"
	"time"
)

type Fetcher interface {
	ListUsers(ctx context.Context) ([]model.User, error)
	ListUserBalances(ctx context.Context) ([]model.UserBalance, error)
	ListTransactions(ctx context.Context) ([]model.Transaction, error)
}

type Service interface {
	ListUserBalances(ctx context.Context) ([]model.UserBalance, error)
	ListTransactions(ctx context.Context) ([]model.Transaction, error)

	FindUsers(ctx context.Context, args *api.Query) ([]model.User, error)
	FindUserByID(ctx context.Context, id string) (*model.User, error)

	FindTransactions(ctx context.Context, args *api.Query) ([]model.Transaction, error)
	FindTransaction(ctx context.Context, id string) (*model.Transaction, error)
}

type UserRepository interface {
	Find(ctx context.Context, args *api.Query) ([]model.User, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByName(ctx context.Context, name string) (*model.User, error)
	Save(ctx context.Context, user model.User) (*model.User, error)
}

type TransactionRepository interface {
	Find(ctx context.Context, args *api.Query) ([]model.Transaction, error)
	FindByID(ctx context.Context, userID string) (*model.Transaction, error)
	Save(ctx context.Context, tx model.Transaction) (*model.Transaction, error)
	FindByTimeAndTotalValueAndSummary(ctx context.Context, time time.Time, value float64, summary string) (*model.Transaction, error)
}
