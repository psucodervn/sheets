package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"time"
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

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type OrderBy string

type Filter string

type Query struct {
	Pagination *Pagination
	OrderBy    OrderBy
	Filter     *Filter
}

func PrepareDB(queryArgs *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if queryArgs == nil {
			return db
		}
		if queryArgs.Filter != nil {
			db = db.Where(*queryArgs.Filter)
		}
		if len(queryArgs.OrderBy) > 0 {
			db = db.Order(string(queryArgs.OrderBy))
		}
		if queryArgs.Pagination != nil {
			if queryArgs.Pagination.Limit > 0 {
				db = db.Limit(queryArgs.Pagination.Limit)
			}
			if queryArgs.Pagination.Offset > 0 {
				db = db.Offset(queryArgs.Pagination.Offset)
			}
		}
		return db
	}
}
