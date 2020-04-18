package balance

import (
	"api/api"
	"api/model"
	"context"
)

var _ Service = &BaseService{}

type BaseService struct {
	fetcher  Fetcher
	userRepo UserRepository
	txRepo   TransactionRepository
}

func (u *BaseService) FindTransaction(ctx context.Context, id string) (*model.Transaction, error) {
	return u.txRepo.FindByID(ctx, id)
}

func (u *BaseService) FindTransactions(ctx context.Context, args *api.Query) ([]model.Transaction, error) {
	return u.txRepo.Find(ctx, args)
}

func (u *BaseService) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

func (u *BaseService) FindUsers(ctx context.Context, args *api.Query) ([]model.User, error) {
	// return u.fetcher.ListUsers(ctx)
	return u.userRepo.Find(ctx, args)
}

func (u *BaseService) ListUserBalances(ctx context.Context) ([]model.UserBalance, error) {
	return u.fetcher.ListUserBalances(ctx)
}

func (u *BaseService) ListTransactions(ctx context.Context) ([]model.Transaction, error) {
	return u.fetcher.ListTransactions(ctx)
}

func NewBaseService(fetcher Fetcher, userRepo UserRepository, txRepo TransactionRepository) *BaseService {
	return &BaseService{
		fetcher:  fetcher,
		userRepo: userRepo,
		txRepo:   txRepo,
	}
}
