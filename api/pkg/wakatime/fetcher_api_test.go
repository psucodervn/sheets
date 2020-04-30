package wakatime

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func TestApiFetcher_FetchTodayCodedTime(t *testing.T) {
	type fields struct {
		apiKey string
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: string("success"), wantErr: false},
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fetcher := newFetcherFromEnv()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetcher.FetchTodayCodedTime(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchTodayCodedTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Today Coded Time: %v\n", got)
		})
	}
}

func TestApiFetcher_FetchLeaderboard(t *testing.T) {
	type fields struct {
		apiKey string
	}
	tests := []struct {
		name    string
		lbID    string
		wantErr bool
	}{
		{name: string("success"), wantErr: false, lbID: "3bd8a420-d89e-43f3-8522-2b7fd12549ba"},
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	fetcher := newFetcherFromEnv()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetcher.FetchLeaderboard(ctx, tt.lbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchLeaderboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ar := make([]string, 0)
			for id, u := range got {
				ar = append(ar, u.Username+":"+string(id))
			}
			t.Log(strings.Join(ar, ","))
		})
	}
}

func newFetcherFromEnv() *ApiFetcher {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	return NewApiFetcher(apiKey)
}
