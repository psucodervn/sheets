package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func DefaultQuery() Query {
	return Query{
		Pagination: Pagination{
			Offset: 0,
			Limit:  100,
		},
	}
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Order struct {
	OrderBy    string `json:"orderBy"`
	Descending bool   `json:"descending"`
}

type Filter struct {
	Condition string        `json:"condition"`
	Args      []interface{} `json:"args"`
}

type Query struct {
	Pagination
	Order
	Filter
}

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}

func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}

type UserClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
