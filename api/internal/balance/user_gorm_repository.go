package balance

import (
	"context"

	"github.com/jinzhu/gorm"

	"api/internal/api"
	"api/oldmodel"
)

var _ UserRepository = &GormPostgresUserRepository{}

type GormPostgresUserRepository struct {
	db *gorm.DB
}

func (r *GormPostgresUserRepository) FindByName(ctx context.Context, name string) (*oldmodel.User, error) {
	var u oldmodel.User
	if err := r.db.Where("name = ?", name).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *GormPostgresUserRepository) Save(ctx context.Context, user oldmodel.User) (*oldmodel.User, error) {
	err := r.db.Save(&user).Error
	return &user, err
}

func (r *GormPostgresUserRepository) FindByID(ctx context.Context, id string) (*oldmodel.User, error) {
	var user oldmodel.User
	err := r.db.First(&user, "id = ?", id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, ErrNotFound
	}
	return &user, err
}

func (r *GormPostgresUserRepository) Find(ctx context.Context, args *api.Query) ([]oldmodel.User, error) {
	users := make([]oldmodel.User, 0)
	err := oldmodel.PrepareDB(args)(r.db).Find(&users).Error
	return users, err
}

func NewGormPostgresUserRepository(db *gorm.DB) *GormPostgresUserRepository {
	return &GormPostgresUserRepository{db: db}
}
