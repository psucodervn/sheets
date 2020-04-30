package balance

import (
	"context"

	"api/internal/api"
	"api/oldmodel"
)

var _ Service = &BaseService{}

type BaseService struct {
	fetcher  Fetcher
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

func (u *BaseService) ListUserBalances(ctx context.Context) ([]oldmodel.UserBalance, error) {
	return u.fetcher.ListUserBalances(ctx)
}

func (u *BaseService) ListTransactions(ctx context.Context) ([]oldmodel.Transaction, error) {
	return u.fetcher.ListTransactions(ctx)
}

func NewBaseService(fetcher Fetcher, userRepo UserRepository, txRepo TransactionRepository) *BaseService {
	return &BaseService{
		fetcher:  fetcher,
		userRepo: userRepo,
		txRepo:   txRepo,
	}
}
