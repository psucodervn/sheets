package model

import (
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type DayOff struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	UserID string    `json:"userId"`
	Date   time.Time `json:"date"`
	Part   string    `json:"part"`
	Note   string    `json:"note"`
}

func (d *DayOff) TableName() string {
	return "days_off"
}

func (d *DayOff) BeforeCreate(db *gorm.DB) error {
	if len(d.ID) == 0 {
		d.ID = xid.New().String()
	}
	return nil
}

type DayOffToCreate struct {
	UserID string    `json:"userId" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
	Part   string    `json:"part" validate:"required"`
	Note   string    `json:"note" validate:""`
}

type DayOffToUpdate struct {
	ID     string    `json:"id" validate:"required"`
	UserID string    `json:"userId" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
	Part   string    `json:"part" validate:"required"`
	Note   string    `json:"note" validate:""`
}
