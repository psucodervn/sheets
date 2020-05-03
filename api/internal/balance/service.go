package balance

import (
	"context"
	"database/sql"
	"sync"

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

	users sync.Map
}

func (s *service) UpdateTransaction(ctx context.Context, id string, txDTO *TransactionDTO) error {
	tx := mapTransactionDTOtoModelTransaction(txDTO)
	tx.ID = id
	_, err := tx.Update(ctx, s.db, boil.Blacklist(model.UserColumns.ID))
	return err
}

func (s *service) DeleteTransaction(ctx context.Context, id string) error {
	tx := &model.Transaction{ID: id}
	n, err := tx.Delete(ctx, s.db, false)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrTransactionNotFound
	}
	return nil
}

func (s *service) AddTransaction(ctx context.Context, tx *model.Transaction) (*model.Transaction, error) {
	err := tx.Insert(ctx, s.db, boil.Infer())
	return tx, err
}

func (s *service) Transactions(ctx context.Context, args api.Query) (model.TransactionSlice, error) {
	return model.Transactions().All(ctx, s.db)
}

func (s *service) User(ctx context.Context, id string) (*model.UserWithBalance, error) {
	return model.UsersWithBalance(model.UserWhere.ID.EQ(id)).One(ctx, s.db)
}

func (s *service) Users(ctx context.Context, args api.Query) ([]model.UserWithBalance, error) {
	return model.UsersWithBalance().All(ctx, s.db)
}

func (s *service) Transaction(ctx context.Context, id string) (*TransactionDTO, error) {
	tx, err := model.FindTransaction(ctx, s.db, id)
	if err == sql.ErrNoRows {
		return nil, ErrTransactionNotFound
	} else if err != nil {
		return nil, err
	}

	txDTO := &TransactionDTO{
		ID:           tx.ID,
		CreatorID:    tx.CreatorID,
		Time:         tx.Time,
		Value:        tx.Value,
		Summary:      tx.Summary,
		Description:  tx.Description,
		Payers:       s.mapChanges(tx.Payers),
		Participants: s.mapChanges(tx.Participants),
		SplitType:    tx.SplitType,
	}
	return txDTO, err
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

func (s *service) mapChanges(changes model.Changes) ChangesDTO {
	changesDTO := make(ChangesDTO, len(changes))
	for i, c := range changes {
		changesDTO[i].ID = c.ID
		changesDTO[i].Value = c.Value
		changesDTO[i].Percent = c.Percent
		changesDTO[i].Name = s.getUserName(c.ID)
	}
	return changesDTO
}

func (s *service) getUserName(id string) string {
	if name, ok := s.users.Load(id); ok {
		return name.(string)
	}
	u, err := model.FindUser(context.Background(), s.db, id, model.UserColumns.Name)
	if err != nil {
		return ""
	}
	s.users.Store(id, u.Name)
	return u.Name
}

func NewService(userRepo UserRepository, txRepo TransactionRepository, db boil.ContextExecutor) *service {
	return &service{
		userRepo: userRepo,
		txRepo:   txRepo,
		db:       db,
	}
}
