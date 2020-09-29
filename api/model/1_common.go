package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Change struct {
	ID      string  `json:"id"`
	Value   float64 `json:"value"`
	Percent float64 `json:"percent,omitempty"`
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

type SplitType uint8

const (
	SplitEqual SplitType = 0
	SplitRatio SplitType = 1
)

type Action string

const (
	ActionCreate Action = "CREATE"
	ActionUpdate Action = "UPDATE"
	ActionRemove Action = "REMOVE"
)

type Meta map[string]interface{}

func (m *Meta) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		return json.Unmarshal(src.([]byte), m)
	default:
		return fmt.Errorf("unsupported type: %v", reflect.TypeOf(src).String())
	}
}

func (m Meta) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type Pager struct {
	Limit  int      `query:"limit"`
	Offset int      `query:"offset"`
	Orders []string `query:"orderBy"`
}

func WithPager(db *gorm.DB, pager Pager) *gorm.DB {
	db = db.Offset(pager.Offset).Limit(pager.Limit)
	for _, ord := range pager.Orders {
		ord = strings.ToLower(strings.TrimSpace(ord))
		if len(ord) == 0 {
			continue
		}
		col, desc := "", false
		if ord[0] == '-' {
			desc = true
			col = ord[1:]
		} else if ord[0] == '+' {
			col = ord[1:]
		} else {
			col = ord
		}
		if len(col) == 0 {
			continue
		}
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: col, Raw: false}, Desc: desc})
	}
	return db
}
