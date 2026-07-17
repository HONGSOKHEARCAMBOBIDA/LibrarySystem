package service

import (
	"context"
	"mysql/config"
	"mysql/constant/apperror"
	"mysql/model"
	"mysql/utils"

	"gorm.io/gorm"
)

type CabinetService interface {
	GetCabinet(ctx context.Context) ([]model.Cabinet, error)
}

type cabinetservice struct {
	db *gorm.DB
}

func NewCabinetService() CabinetService {
	return &cabinetservice{
		db: config.DB,
	}
}

func (s *cabinetservice) GetCabinet(ctx context.Context) ([]model.Cabinet, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.DefaultQueryTimeout)
	defer cancel()

	var cabinets []model.Cabinet
	err := s.db.WithContext(ctx).Find(&cabinets).Error
	if err != nil {
		return nil, apperror.Internal("failed to fetch cabinets", err)
	}
	return cabinets, nil
}
