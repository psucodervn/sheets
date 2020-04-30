package balance

import (
	"context"
	"time"

	"api/internal/api"
	"api/oldmodel"
)

type Fetcher interface {
	ListUsers(ctx context.Context) ([]oldmodel.User, error)
	ListUserBalances(ctx context.Context) ([]oldmodel.UserBalance, error)
	ListTransactions(ctx context.Context) ([]oldmodel.Transaction, error)
}

type Service interface {
	FindUsers(ctx context.Context, args *api.Query) ([]oldmodel.User, error)
	FindUserByID(ctx context.Context, id string) (*oldmodel.User, error)

	FindTransactions(ctx context.Context, args *api.Query) ([]oldmodel.Transaction, error)
	FindTransaction(ctx context.Context, id string) (*oldmodel.Transaction, error)
}

type UserRepository interface {
	Find(ctx context.Context, args *api.Query) ([]oldmodel.User, error)
	FindByID(ctx context.Context, id string) (*oldmodel.User, error)
	FindByName(ctx context.Context, name string) (*oldmodel.User, error)
	Save(ctx context.Context, user oldmodel.User) (*oldmodel.User, error)
}

type TransactionRepository interface {
	Find(ctx context.Context, args *api.Query) ([]oldmodel.Transaction, error)
	FindByID(ctx context.Context, userID string) (*oldmodel.Transaction, error)
	Save(ctx context.Context, tx oldmodel.Transaction) (*oldmodel.Transaction, error)
	FindByTimeAndTotalValueAndSummary(ctx context.Context, time time.Time, value float64, summary string) (*oldmodel.Transaction, error)
}
