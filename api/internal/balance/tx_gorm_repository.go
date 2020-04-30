package balance

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"api/internal/api"
	"api/oldmodel"
)

var _ TransactionRepository = &GormPostgresTransactionRepository{}

type GormPostgresTransactionRepository struct {
	db *gorm.DB
}

func (r *GormPostgresTransactionRepository) FindByTimeAndTotalValueAndSummary(ctx context.Context, time time.Time, value float64, summary string) (*oldmodel.Transaction, error) {
	var tx oldmodel.Transaction
	err := r.db.First(&tx, "time = ? AND total_value = ? AND summary = ?", time, value, summary).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, ErrNotFound
	}
	return &tx, err
}

func (r *GormPostgresTransactionRepository) Save(ctx context.Context, tx oldmodel.Transaction) (*oldmodel.Transaction, error) {
	err := r.db.Save(&tx).Error
	return &tx, err
}

func (r *GormPostgresTransactionRepository) Find(ctx context.Context, args *api.Query) ([]oldmodel.Transaction, error) {
	txs := make([]oldmodel.Transaction, 0)
	err := oldmodel.PrepareDB(args)(r.db).Find(&txs).Error
	return txs, err
}

func (r *GormPostgresTransactionRepository) FindByID(ctx context.Context, txID string) (*oldmodel.Transaction, error) {
	var tx oldmodel.Transaction
	err := r.db.First(&tx, "id = ?", txID).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, ErrNotFound
	}
	return &tx, err
}

func NewGormPostgresTransactionRepository(db *gorm.DB) *GormPostgresTransactionRepository {
	return &GormPostgresTransactionRepository{db: db}
}
