package balance

import (
	"context"

	"api/internal/api"
	"api/model"
)

type Service interface {
	Transaction(ctx context.Context, id string) (*TransactionDTO, error)
	Users(ctx context.Context, args api.Query) ([]model.UserWithBalance, error)
	User(ctx context.Context, id string) (*model.UserWithBalance, error)
	Transactions(ctx context.Context, args api.Query) ([]TransactionDTO, error)
	AddTransaction(ctx context.Context, tx *model.Transaction, user *model.User) (*model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string, user *model.User) error
	UpdateTransaction(ctx context.Context, id string, tx *TransactionDTO, user *model.User) error
}
