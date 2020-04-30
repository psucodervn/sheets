package point

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"api/oldmodel"
	"api/pkg/wakatime"
)

type ReportService interface {
	GetReport(ctx context.Context, from time.Time, to time.Time) ([]oldmodel.UserPoint, error)
}

var _ ReportService = &BaseReportService{}

type BaseReportService struct {
	pointSvc          Service
	wakaSvc           wakatime.Fetcher
	wakaLeaderboardID string
	mapNameToWakaID   map[string]string
}

func NewBaseReportService(pointSvc Service, wakaSvc wakatime.Fetcher, wakaLeaderboardID string, mapNameToWakaID map[string]string) *BaseReportService {
	return &BaseReportService{pointSvc: pointSvc, wakaSvc: wakaSvc, wakaLeaderboardID: wakaLeaderboardID, mapNameToWakaID: mapNameToWakaID}
}

func (s *BaseReportService) GetReport(ctx context.Context, from time.Time, to time.Time) ([]oldmodel.UserPoint, error) {
	users, err := s.pointSvc.WorkingIssues(ctx, from, to)
	if err != nil {
		return nil, err
	}
	wakaUsers, err := s.wakaSvc.FetchLeaderboard(ctx, s.wakaLeaderboardID)
	if err != nil {
		return nil, err
	}

	for i := range users {
		id := wakatime.UserID(s.mapNameToWakaID[users[i].Name])
		if len(id) == 0 {
			log.Warn().Str("jira_name", users[i].Name).Msg("not found waka user")
			continue
		}
		users[i].WakatimeHuman = wakaUsers[id].Human
		users[i].WakatimeSeconds = wakaUsers[id].Seconds
	}
	return users, nil
}
