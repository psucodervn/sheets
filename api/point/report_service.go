package point

import (
	"context"
	"time"
)

type ReportService interface {
	GetReport(ctx context.Context, from time.Time, to time.Time) (*GetReportResponse, error)
}

var _ ReportService = &BaseReportService{}

type BaseReportService struct {
}

func (s *BaseReportService) GetReport(ctx context.Context, from time.Time, to time.Time) (*GetReportResponse, error) {
	panic("implement me")
}
