package balance

import (
	"api/model"
	"context"
)

var _ Service = &BaseService{}

type BaseService struct {
	fetcher  Fetcher
	userRepo UserRepository
}

func (u *BaseService) FindByID(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

func (u *BaseService) Find(ctx context.Context, args *model.Query) ([]model.User, error) {
	// return u.fetcher.ListUsers(ctx)
	return u.userRepo.Find(ctx, args)
}

func (u *BaseService) ListUserBalances(ctx context.Context) ([]model.UserBalance, error) {
	return u.fetcher.ListUserBalances(ctx)
}

func (u *BaseService) ListTransactions(ctx context.Context) ([]model.Transaction, error) {
	return u.fetcher.ListTransactions(ctx)
}

func NewBaseService(fetcher Fetcher, userRepo UserRepository) *BaseService {
	return &BaseService{
		fetcher:  fetcher,
		userRepo: userRepo,
	}
}
