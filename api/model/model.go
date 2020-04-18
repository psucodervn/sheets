package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"

	"api/api"
)

type Model struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("ID", xid.New().String()); err != nil {
		return err
	}
	if err := scope.SetColumn("CreatedAt", time.Now()); err != nil {
		return err
	}
	if err := scope.SetColumn("UpdatedAt", time.Now()); err != nil {
		return err
	}
	return nil
}

func PrepareDB(queryArgs *api.Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if queryArgs == nil {
			return db
		}
		if len(queryArgs.Filter) > 0 {
			db = db.Where(queryArgs.Filter)
		}
		if len(queryArgs.OrderBy) > 0 {
			if queryArgs.Descending {
				db = db.Order(queryArgs.OrderBy + " DESC")
			} else {
				db = db.Order(queryArgs.OrderBy)
			}
		}
		if queryArgs.Pagination.Limit > 0 {
			db = db.Limit(queryArgs.Pagination.Limit)
		}
		if queryArgs.Pagination.Offset > 0 {
			db = db.Offset(queryArgs.Pagination.Offset)
		}
		return db
	}
}
