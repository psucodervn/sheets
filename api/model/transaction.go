package model

import "time"

type Transaction struct {
	Description string            `json:"description"`
	Value       float64           `json:"value"`
	Time        time.Time         `json:"time"`
	Senders     []UserTransaction `json:"senders"`
	Receivers   []UserTransaction `json:"receivers"`
}

type UserTransaction struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}
