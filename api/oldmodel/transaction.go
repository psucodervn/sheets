package oldmodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	Model
	Summary     string           `json:"summary"`
	Description string           `json:"description"`
	TotalValue  float64          `json:"totalValue"`
	Time        time.Time        `json:"time" gorm:"not null"`
	Senders     UserTransactions `json:"senders" gorm:"not null;type:jsonb;default:'[]'"`
	Receivers   UserTransactions `json:"receivers" gorm:"not null;type:jsonb;default:'[]'"`
}

type UserTransaction struct {
	ID   string  `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
	Val  float64 `json:"value,omitempty"`
}

type UserTransactions []UserTransaction

func (u UserTransactions) Value() (driver.Value, error) {
	if u == nil {
		return json.Marshal(UserTransactions{})
	}
	return json.Marshal(u)
}

func (u *UserTransactions) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", src))
	}

	return json.Unmarshal(bytes, u)
}
