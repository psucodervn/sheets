package model

type UserWithStar struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Stars float64 `json:"stars"`
}

type UserStars []UserWithStar
