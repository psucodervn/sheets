package balance

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"api/internal/api"
	"api/model"
	"api/oldmodel"
)

var _ Service = &service{}

type service struct {
	userRepo UserRepository
	txRepo   TransactionRepository
	db       boil.ContextExecutor
}

func (s *service) User(ctx context.Context, id string) (*model.UserWithBalance, error) {
	return model.UsersWithBalance(model.UserWhere.ID.EQ(id)).One(ctx, s.db)
}

func (s *service) Users(ctx context.Context, args api.Query) ([]model.UserWithBalance, error) {
	return model.UsersWithBalance().All(ctx, s.db)
}

func (s *service) FindTransaction(ctx context.Context, id string) (*oldmodel.Transaction, error) {
	return s.txRepo.FindByID(ctx, id)
}

func (s *service) FindTransactions(ctx context.Context, args *api.Query) ([]oldmodel.Transaction, error) {
	return s.txRepo.Find(ctx, args)
}

func (s *service) FindUserByID(ctx context.Context, id string) (*oldmodel.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *service) FindUsers(ctx context.Context, args *api.Query) ([]oldmodel.User, error) {
	// return u.fetcher.ListUsers(ctx)
	return s.userRepo.Find(ctx, args)
}

func NewService(userRepo UserRepository, txRepo TransactionRepository, db boil.ContextExecutor) *service {
	return &service{
		userRepo: userRepo,
		txRepo:   txRepo,
		db:       db,
	}
}
