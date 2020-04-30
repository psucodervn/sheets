package balance

import (
	"context"

	"api/internal/api"
	"api/oldmodel"
)

var _ Service = &BaseService{}

type BaseService struct {
	userRepo UserRepository
	txRepo   TransactionRepository
}

func (u *BaseService) FindTransaction(ctx context.Context, id string) (*oldmodel.Transaction, error) {
	return u.txRepo.FindByID(ctx, id)
}

func (u *BaseService) FindTransactions(ctx context.Context, args *api.Query) ([]oldmodel.Transaction, error) {
	return u.txRepo.Find(ctx, args)
}

func (u *BaseService) FindUserByID(ctx context.Context, id string) (*oldmodel.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

func (u *BaseService) FindUsers(ctx context.Context, args *api.Query) ([]oldmodel.User, error) {
	// return u.fetcher.ListUsers(ctx)
	return u.userRepo.Find(ctx, args)
}

func NewBaseService(userRepo UserRepository, txRepo TransactionRepository) *BaseService {
	return &BaseService{
		userRepo: userRepo,
		txRepo:   txRepo,
	}
}
