package user

import (
	"context"
	"database/sql"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"api/model"
)

type Service struct {
	db boil.ContextExecutor
}

func NewService(db boil.ContextExecutor) *Service {
	return &Service{db: db}
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := model.Users(model.UserWhere.Email.EQ(null.StringFrom(email))).One(ctx, s.db)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	return u, err
}

func (s *Service) FindByAuthProvider(ctx context.Context, provider string, id string) (*model.User, error) {
	ai, err := model.AuthIdentities(
		model.AuthIdentityWhere.ID.EQ(id),
		model.AuthIdentityWhere.Provider.EQ(provider),
		qm.Load(qm.Rels(model.AuthIdentityRels.User)),
	).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAuthNotFound
		}
		return nil, err
	}
	return ai.R.User, err
}
