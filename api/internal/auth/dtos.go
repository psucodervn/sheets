package auth

import (
	"api/model"
)

type GoogleUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	SheetName string `json:"sheetName"`
	JiraName  string `json:"jiraName"`
}

func (u *User) FromModel(user *model.User) *User {
	u.ID = user.ID
	u.Email = user.Email.String
	u.Name = user.Name
	u.SheetName = user.SheetName.String
	u.JiraName = user.JiraName.String
	return u
}

type Token struct {
	AccessToken string `json:"accessToken"`
}
