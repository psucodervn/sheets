package balance

import (
	"context"
	"time"

	"api/internal/api"
	"api/model"
	"api/oldmodel"
)

type Fetcher interface {
	ListUsers(ctx context.Context) ([]oldmodel.User, error)
	ListUserBalances(ctx context.Context) ([]oldmodel.UserBalance, error)
	ListTransactions(ctx context.Context) ([]oldmodel.Transaction, error)
}

type Service interface {
	Transaction(ctx context.Context, id string) (*TransactionDTO, error)
	Users(ctx context.Context, args api.Query) ([]model.UserWithBalance, error)
	User(ctx context.Context, id string) (*model.UserWithBalance, error)
	Transactions(ctx context.Context, args api.Query) ([]TransactionDTO, error)
	AddTransaction(ctx context.Context, tx *model.Transaction, user *model.User) (*model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string, user *model.User) error
	UpdateTransaction(ctx context.Context, id string, tx *TransactionDTO, user *model.User) error
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
