package model

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Balance struct {
	Value float64 `json:"value"`
}

type UserBalance struct {
	User    User    `json:"user"`
	Balance Balance `json:"balance"`
}
