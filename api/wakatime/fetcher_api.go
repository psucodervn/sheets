package wakatime

import (
	"context"
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

var _ Fetcher = &ApiFetcher{}

type ApiFetcher struct {
	client *resty.Client
}

func (f *ApiFetcher) FetchLeaderboard(ctx context.Context, leaderboardID string) (Users, error) {
	resp, err := f.client.R().
		SetResult(&LeaderboardLeaders{}).
		SetContext(ctx).
		Get("https://wakatime.com/api/v1/users/current/leaderboards/" + leaderboardID)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.Status())
	}

	res := resp.Result().(*LeaderboardLeaders)
	m := make(Users)
	for _, u := range res.Data {
		m[UserID(u.User.ID)] = UserTime{
			Username:    u.User.Username,
			DisplayName: u.User.DisplayName,
			Human:       u.RunningTotal.HumanReadableTotal,
			Seconds:     u.RunningTotal.TotalSeconds,
		}
	}
	return m, nil
}

type summariesResponse struct {
	Data []SummaryData `json:"data"`
}

func (f *ApiFetcher) FetchTodayCodedTime(ctx context.Context) (time.Duration, error) {
	start := time.Now().Format("2006-01-02")
	end := start

	resp, err := f.client.R().
		SetQueryParam("start", start).
		SetQueryParam("end", end).
		SetResult(&summariesResponse{}).
		SetContext(ctx).
		Get("https://wakatime.com/api/v1/users/current/summaries")
	if err != nil {
		return 0, err
	}
	if resp.IsError() {
		return 0, errors.New(resp.Status())
	}

	res := resp.Result().(*summariesResponse)
	if len(res.Data) == 0 {
		return 0, errors.New("invalid response")
	}

	ms := time.Duration(res.Data[0].GrandTotal.TotalSeconds * 1000)
	return ms * time.Millisecond, nil
}

func NewApiFetcher(apiKey string) *ApiFetcher {
	return &ApiFetcher{
		client: resty.New().SetBasicAuth(apiKey, ""),
	}
}
