package telegram

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/dustin/go-humanize"

	"api/model"
)

var (
	LocalZone = time.FixedZone("UTC+7", 7*60*60)
)

type Users []model.UserWithBalance

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Balance > u[j].Balance
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func marshalUser(user *model.UserWithBalance) string {
	return fmt.Sprintf("%s: %s (vnđ)", user.Name, humanize.Comma(int64(user.Balance)))
}

func marshalTopUsers(users Users, top int) string {
	sort.Sort(users)
	bf := bytes.NewBuffer(nil)
	if top < math.MaxInt32 {
		bf.WriteString(fmt.Sprintf("Top ±%d users:", top))
	} else {
		bf.WriteString("All users:")
	}
	dot := false
	for i, u := range users {
		if i >= top && i < len(users)-top {
			if !dot {
				dot = true
				bf.WriteString("\n......")
			}
			continue
		}
		bf.WriteString(fmt.Sprintf("\n%d. ", i+1))
		bf.WriteString(marshalUser(&u))
	}
	return bf.String()
}

type user struct {
	ID string
}

func newUser(id string) *user {
	return &user{ID: id}
}

func (u user) Recipient() string {
	return u.ID
}
