package balance

import (
	"context"
	"database/sql"
	"sync"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/sync/errgroup"

	"api/model"
	"api/oldmodel"
)

type Importer interface {
	Run() error
}

var _ Importer = &importer{}

type importer struct {
	fetcher Fetcher
	db      boil.ContextExecutor

	users sync.Map
}

func NewImporter(fetcher Fetcher, db boil.ContextExecutor) *importer {
	return &importer{fetcher: fetcher, db: db}
}

func (imp *importer) Run() error {
	ctx := context.Background()

	users, txs, err := fetchUserAndTx(ctx, imp.fetcher)
	if err != nil {
		return err
	}

	if err := imp.saveUsers(ctx, users); err != nil {
		return err
	}

	if err := imp.saveTransactions(ctx, txs); err != nil {
		return err
	}

	return nil
}

func (imp *importer) saveUsers(ctx context.Context, users []oldmodel.User) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, u := range users {
		u := &model.User{
			Name:      u.Name,
			SheetName: null.StringFrom(u.Name),
		}
		eg.Go(func() error {
			if u, err := model.Users(model.UserWhere.SheetName.EQ(null.StringFrom(u.Name))).One(ctx, imp.db); err == nil {
				imp.users.Store(u.SheetName.String, u.ID)
				return nil
			}
			err := u.Insert(ctx, imp.db, boil.Infer())
			if err != nil {
				return err
			}
			imp.users.Store(u.SheetName.String, u.ID)
			return nil
		})
	}
	return eg.Wait()
}

func (imp *importer) saveTransactions(ctx context.Context, txs []oldmodel.Transaction) error {
	fromUserTransactions := func(txs oldmodel.UserTransactions) model.Changes {
		changes := make(model.Changes, len(txs))
		for i := range txs {
			id, _ := imp.users.Load(txs[i].Name)
			changes[i].ID = id.(string)
			changes[i].Value = txs[i].Val
			changes[i].Percent = txs[i].Percent
		}
		return changes
	}

	eg, ctx := errgroup.WithContext(ctx)
	for i := range txs {
		i := i
		eg.Go(func() error {
			tx, err := model.Transactions(
				model.TransactionWhere.Time.EQ(txs[i].Time),
				model.TransactionWhere.Summary.EQ(txs[i].Summary),
				model.TransactionWhere.Value.EQ(txs[i].TotalValue),
			).One(ctx, imp.db)
			if err == sql.ErrNoRows {
				tx = &model.Transaction{
					Time:    txs[i].Time,
					Value:   txs[i].TotalValue,
					Summary: txs[i].Summary,
				}
			} else if err != nil {
				return err
			}
			tx.SplitType = model.SplitRatio
			tx.Description = null.StringFrom(txs[i].Description)
			tx.Payers = fromUserTransactions(txs[i].Senders)
			tx.Participants = fromUserTransactions(txs[i].Receivers)
			return tx.Upsert(ctx, imp.db, true, []string{}, boil.Infer(), boil.Infer())
		})
	}
	return eg.Wait()
}
