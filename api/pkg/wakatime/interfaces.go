package wakatime

import (
	"context"
	"time"
)

type Fetcher interface {
	FetchTodayCodedTime(ctx context.Context) (time.Duration, error)
	FetchLeaderboard(ctx context.Context, leaderboardID string) (Users, error)
}
