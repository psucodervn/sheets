package balance

import (
	"api/model"
	"context"
	"github.com/jinzhu/gorm"
)

var _ TransactionRepository = &PostgresTransactionRepository{}

type PostgresTransactionRepository struct {
	db *gorm.DB
}

func (r *PostgresTransactionRepository) Save(ctx context.Context, tx model.Transaction) (*model.Transaction, error) {
	err := r.db.Save(&tx).Error
	return &tx, err
}

func (r *PostgresTransactionRepository) Find(ctx context.Context, args *model.Query) ([]model.Transaction, error) {
	txs := make([]model.Transaction, 0)
	err := model.PrepareDB(args)(r.db).Find(&txs).Error
	return txs, err
}

func (r *PostgresTransactionRepository) FindByID(ctx context.Context, txID string) (*model.Transaction, error) {
	var tx model.Transaction
	err := r.db.First(&tx, "id = ?", txID).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, ErrNotFound
	}
	return &tx, err
}

func NewPostgresTransactionRepository(db *gorm.DB) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{db: db}
}
