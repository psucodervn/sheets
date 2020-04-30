package balance

import (
	"context"

	"golang.org/x/sync/errgroup"

	"api/oldmodel"
)

type OldImporter struct {
	fetcher  Fetcher
	userRepo UserRepository
	txRepo   TransactionRepository
}

func NewOldImporter(fetcher Fetcher, userRepo UserRepository, txRepo TransactionRepository) *OldImporter {
	return &OldImporter{fetcher: fetcher, userRepo: userRepo, txRepo: txRepo}
}

func (s *OldImporter) Run() error {
	ctx := context.Background()

	users, userTxs, err := fetchUserAndTx(ctx, s.fetcher)
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

func fetchUserAndTx(ctx context.Context, fetcher Fetcher) ([]oldmodel.User, []oldmodel.Transaction, error) {
	wg, ctx := errgroup.WithContext(ctx)
	var users []oldmodel.User
	var userTxs []oldmodel.Transaction

	wg.Go(func() error {
		var err error
		users, err = fetcher.ListUsers(ctx)
		return err
	})
	wg.Go(func() error {
		var err error
		userTxs, err = fetcher.ListTransactions(ctx)
		return err
	})

	if err := wg.Wait(); err != nil {
		return nil, nil, err
	}
	return users, userTxs, nil
}

func (s *OldImporter) saveUsers(ctx context.Context, users []oldmodel.User) error {
	wg := new(errgroup.Group)
	for i := range users {
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

func (s *OldImporter) saveTransactions(ctx context.Context, txs []oldmodel.Transaction, users []oldmodel.User) error {
	mu := make(map[string]oldmodel.User)
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
			if _, err := s.txRepo.FindByTimeAndTotalValueAndSummary(ctx, tx.Time, tx.TotalValue, tx.Summary); err == nil {
				// duplicate, ignore
				return nil
			} else if err != ErrNotFound {
				return err
			}
			_, err := s.txRepo.Save(ctx, tx)
			return err
		})
	}
	return wg.Wait()
}
