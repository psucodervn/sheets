package repository

import (
	"context"
	"testing"
)

func TestApiFetcher_List(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		debug   bool
	}{
		{wantErr: false},
	}
	f := NewApiFetcherFromEnv()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.ListUserBalances(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUserBalances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("ListUserBalances() return 0 items")
			}
			if tt.debug {
				t.Logf("ListUserBalances() returns: %+v", got)
			}
		})
	}
}

func TestApiFetcher_ListTransactions(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		debug   bool
	}{
		{wantErr: false, debug: true},
	}
	f := NewApiFetcherFromEnv()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.ListTransactions(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("ListTransactions() return 0 items")
			}
			if tt.debug {
				t.Logf("ListTransactions() returns %v items:\n%+v", len(got), got)
			}
		})
	}
}

func TestApiFetcher_ListUsers(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		debug   bool
	}{
		{wantErr: false, debug: true},
	}
	f := NewApiFetcherFromEnv()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.ListUsers(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("ListUsers() return 0 items")
			}
			if tt.debug {
				t.Logf("ListUsers() returns %v items:\n%+v", len(got), got)
			}
		})
	}
}
