package model

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
)

type UserWithBalance struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Email      null.String `json:"email"`
	Balance    float64     `boil:"balance" json:"balance"`
	TelegramID null.String `json:"telegramID"`
}

type userWithBalanceQuery struct {
	*queries.Query
}

// UsersWithBalance returns userWithBalanceQuery
func UsersWithBalance(mods ...qm.QueryMod) userWithBalanceQuery {
	mods = append(mods,
		qm.From(TableNames.Users), qmhelper.WhereIsNull(UserColumns.DeletedAt),
		qm.LeftOuterJoin(`balance b ON b.user_id = `+UserColumns.ID),
		qm.Select("*", "b.value as balance"),
	)
	return userWithBalanceQuery{NewQuery(mods...)}
}

func (q userWithBalanceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserWithBalance, error) {
	var user UserWithBalance
	err := q.Bind(ctx, exec, &user)
	return &user, err
}

func (q userWithBalanceQuery) All(ctx context.Context, exec boil.ContextExecutor) ([]UserWithBalance, error) {
	var users []UserWithBalance
	err := q.Bind(ctx, exec, &users)
	return users, err
}
