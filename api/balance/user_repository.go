package balance

import (
	"context"

	"github.com/jinzhu/gorm"

	"api/api"
	"api/model"
)

var _ UserRepository = &PostgresUserRepository{}

type PostgresUserRepository struct {
	db *gorm.DB
}

func (r *PostgresUserRepository) FindByName(ctx context.Context, name string) (*model.User, error) {
	var u model.User
	if err := r.db.Where("name = ?", name).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *PostgresUserRepository) Save(ctx context.Context, user model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	return &user, err
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ?", id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, ErrNotFound
	}
	return &user, err
}

func (r *PostgresUserRepository) Find(ctx context.Context, args *api.Query) ([]model.User, error) {
	users := make([]model.User, 0)
	err := model.PrepareDB(args)(r.db).Find(&users).Error
	return users, err
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}
