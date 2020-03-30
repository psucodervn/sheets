package importer

import (
	"api/balance"
	"api/model"
	"context"
	"golang.org/x/sync/errgroup"
)

type Importer struct {
	fetcher  balance.Fetcher
	userRepo balance.UserRepository
	txRepo   balance.TransactionRepository
}

func NewImporter(fetcher balance.Fetcher, userRepo balance.UserRepository, txRepo balance.TransactionRepository) *Importer {
	return &Importer{fetcher: fetcher, userRepo: userRepo, txRepo: txRepo}
}

func (s *Importer) Run() error {
	ctx := context.Background()

	users, userTxs, err := s.fetchUserAndTx(ctx)
	if err != nil {
		return err
	}

	if err := s.saveUsers(ctx, users); err != nil {
		return err
	}

	if err := s.saveTransactions(ctx, userTxs, users); err != nil {
		return err
	}

	return nil
}

func (s *Importer) fetchUserAndTx(ctx context.Context) ([]model.User, []model.Transaction, error) {
	wg, ctx := errgroup.WithContext(ctx)
	var users []model.User
	var userTxs []model.Transaction

	wg.Go(func() error {
		var err error
		users, err = s.fetcher.ListUsers(ctx)
		return err
	})
	wg.Go(func() error {
		var err error
		userTxs, err = s.fetcher.ListTransactions(ctx)
		return err
	})

	if err := wg.Wait(); err != nil {
		return nil, nil, err
	}
	return users, userTxs, nil
}

func (s *Importer) saveUsers(ctx context.Context, users []model.User) error {
	wg := new(errgroup.Group)
	for i, _ := range users {
		i := i
		wg.Go(func() error {
			if u, _ := s.userRepo.FindByName(ctx, users[i].Name); u != nil {
				users[i].ID = u.ID
			}
			newUser, err := s.userRepo.Save(ctx, users[i])
			if err != nil {
				return err
			}
			users[i] = *newUser
			return nil
		})
	}
	return wg.Wait()
}

func (s *Importer) saveTransactions(ctx context.Context, txs []model.Transaction, users []model.User) error {
	mu := make(map[string]model.User)
	for _, u := range users {
		mu[u.Name] = u
	}

	wg := new(errgroup.Group)
	for i := range txs {
		tx := txs[i]
		wg.Go(func() error {
			tx.Summary = tx.Description
			for j := range tx.Senders {
				tx.Senders[j].ID = mu[tx.Senders[j].Name].ID
			}
			for j := range tx.Receivers {
				tx.Receivers[j].ID = mu[tx.Receivers[j].Name].ID
			}
			_, err := s.txRepo.Save(ctx, tx)
			return err
		})
	}
	return wg.Wait()
}
