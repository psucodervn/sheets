package balance

import (
	"github.com/volatiletech/null/v8"
)

type UserWithBalance struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Email   null.String `json:"email"`
	Balance float64     `boil:"balance" json:"balance"`
}
