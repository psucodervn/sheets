package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type User struct {
	Model
	Name    string   `json:"name" gorm:"not null"`
	Email   string   `json:"email"`
	Balance Balance  `json:"balance" gorm:"type:jsonb;not null;default:'{}'"`
	Info    UserInfo `json:"info" gorm:"type:jsonb;not null;default:'{}'"`
}

type Balance struct {
	Val float64 `json:"value"`
}

func (b *Balance) Scan(src interface{}) error {
	return scanBytes(src, b)
}

func (b Balance) Value() (driver.Value, error) {
	return json.Marshal(b)
}

type UserInfo struct {
	FullName string `json:"email"`
}

func (b *UserInfo) Scan(src interface{}) error {
	return scanBytes(src, b)
}

func (b UserInfo) Value() (driver.Value, error) {
	return json.Marshal(b)
}

type UserBalance struct {
	User    User    `json:"user"`
	Balance Balance `json:"balance"`
}

func scanBytes(src interface{}, b interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", src))
	}
	return json.Unmarshal(bytes, b)
}
