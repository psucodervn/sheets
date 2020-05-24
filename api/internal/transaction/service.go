package transaction

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Service struct {
	db boil.ContextExecutor
}

func NewService(db boil.ContextExecutor) *Service {
	return &Service{db: db}
}
