package model

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	AddTransactionHook(boil.AfterInsertHook, refreshBalance)
	AddTransactionHook(boil.AfterUpdateHook, refreshBalance)
	AddTransactionHook(boil.AfterUpsertHook, refreshBalance)
	AddTransactionHook(boil.AfterDeleteHook, refreshBalance)
}

func refreshBalance(ctx context.Context, db boil.ContextExecutor, tx *Transaction) error {
	_, err := db.ExecContext(ctx, `REFRESH MATERIALIZED VIEW balance`)
	return err
}
