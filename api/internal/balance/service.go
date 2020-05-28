package balance

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"api/internal/api"
	"api/model"
	sheet "api/proto"
)

var _ Service = &service{}

type service struct {
	db                 *sql.DB
	notificationClient sheet.NotificationServiceClient

	users sync.Map
}

func NewService(db *sql.DB, notificationClient sheet.NotificationServiceClient) *service {
	return &service{db: db, notificationClient: notificationClient}
}

func (s *service) UpdateTransaction(ctx context.Context, id string, txDTO *TransactionDTO, user *model.User) (err error) {
	tx := mapTransactionDTOtoModelTransaction(txDTO)
	tx.ID = id

	txn, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = txn.Rollback()
		}
	}()

	oldTx, err := model.FindTransaction(ctx, txn, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrTransactionNotFound
		}
		return err
	}

	if _, err = tx.Update(ctx, txn, boil.Blacklist(model.UserColumns.ID)); err != nil {
		return err
	}

	txLog := &model.TransactionLog{
		TransactionID: tx.ID,
		ActorID:       user.ID,
		Action:        model.ActionUpdate,
		Time:          time.Now(),
		Meta: model.Meta{
			"username": user.Name,
			"oldTx":    oldTx,
			"newTx":    tx,
		},
	}
	if err = txLog.Insert(ctx, txn, boil.Infer()); err != nil {
		return err
	}

	_, err = s.notificationClient.NotifyTransaction(ctx, &sheet.NotifyTransactionRequest{
		TransactionId: tx.ID,
	})
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (s *service) DeleteTransaction(ctx context.Context, id string, user *model.User) error {
	txn, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	tx := &model.Transaction{ID: id}
	n, err := tx.Delete(ctx, txn, false)
	if err != nil {
		_ = txn.Rollback()
		return err
	}
	if n == 0 {
		_ = txn.Rollback()
		return ErrTransactionNotFound
	}

	txLog := &model.TransactionLog{
		TransactionID: tx.ID,
		ActorID:       user.ID,
		Action:        model.ActionRemove,
		Time:          time.Now(),
		Meta:          model.Meta{"username": user.Name},
	}
	if err = txLog.Insert(ctx, txn, boil.Infer()); err != nil {
		_ = txn.Rollback()
		return err
	}
	return txn.Commit()
}

func (s *service) AddTransaction(ctx context.Context, tx *model.Transaction, user *model.User) (*model.Transaction, error) {
	txn, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if err = tx.Insert(ctx, txn, boil.Infer()); err != nil {
		_ = txn.Rollback()
		return nil, err
	}
	txLog := &model.TransactionLog{
		TransactionID: tx.ID,
		ActorID:       user.ID,
		Action:        model.ActionCreate,
		Time:          time.Now(),
		Meta:          model.Meta{"username": user.Name},
	}
	if err = txLog.Insert(ctx, txn, boil.Infer()); err != nil {
		_ = txn.Rollback()
		return nil, err
	}
	err = txn.Commit()
	return tx, err
}

func (s *service) Transactions(ctx context.Context, args api.Query) ([]TransactionDTO, error) {
	txs, err := model.Transactions(qm.OrderBy(model.TransactionColumns.ID+" DESC")).All(ctx, s.db)
	if err != nil || len(txs) == 0 {
		return []TransactionDTO{}, err
	}

	rows := make([]TransactionDTO, len(txs))
	for i, tx := range txs {
		rows[i] = *s.mapModelTransactionToDTO(tx)
	}
	return rows, nil
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
	res := s.db.QueryRow(`SELECT name FROM `+model.TableNames.Users+` WHERE id = $1`, id)
	var name string
	if err := res.Scan(&name); err != nil {
		return ""
	}
	s.users.Store(id, name)
	return name
}

func (s *service) mapModelTransactionToDTO(tx *model.Transaction) *TransactionDTO {
	return &TransactionDTO{
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
}
