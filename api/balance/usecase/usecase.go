package usecase

import (
	"api/balance"
	"api/models"
	"context"
)

type UseCase struct {
	fetcher balance.Fetcher
}

func (u *UseCase) ListUsers(ctx context.Context) ([]models.User, error) {
	return u.fetcher.ListUsers(ctx)
}

func (u *UseCase) ListUserBalances(ctx context.Context) ([]models.UserBalance, error) {
	return u.fetcher.ListUserBalances(ctx)
}

func (u *UseCase) ListTransactions(ctx context.Context) ([]models.Transaction, error) {
	return u.fetcher.ListTransactions(ctx)
}

func NewUseCase(fetcher balance.Fetcher) *UseCase {
	return &UseCase{fetcher: fetcher}
}
