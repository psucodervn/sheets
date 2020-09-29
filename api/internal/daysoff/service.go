package daysoff

import (
	"context"

	"gorm.io/gorm"

	"api/model"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(ctx context.Context, req model.DayOffToCreate) (*model.DayOff, error) {
	do := &model.DayOff{
		UserID: req.UserID,
		Date:   req.Date,
		Part:   req.Part,
		Note:   req.Note,
	}
	if err := s.db.Create(do).Error; err != nil {
		return nil, err
	}
	return do, nil
}

func (s *Service) List(ctx context.Context, pager model.Pager) ([]model.DayOff, error) {
	var days []model.DayOff
	err := model.WithPager(s.db, pager).Order("date DESC").Find(&days).Error
	if err != nil {
		return nil, err
	}
	return days, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	do := &model.DayOff{ID: id}
	return s.db.Delete(do).Error
}
