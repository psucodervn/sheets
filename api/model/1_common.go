package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type Change struct {
	ID    string  `json:"id"`
	Value float64 `json:"value"`
}

type Changes []Change

func (c *Changes) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		return json.Unmarshal(src.([]byte), c)
	default:
		return fmt.Errorf("unsupported type: %v", reflect.TypeOf(src).String())
	}
}

func (c Changes) Value() (driver.Value, error) {
	return json.Marshal(c)
}
