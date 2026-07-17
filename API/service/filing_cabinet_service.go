package service

import (
	"context"
	"mysql/config"
	"mysql/constant/apperror"
	"mysql/model"
	"mysql/utils"

	"gorm.io/gorm"
)

type FilingCabinetService interface {
	GetFilingCabinet(ctx context.Context, cabinetID int) ([]model.FilingCabinet, error)
}

type filingcabinetservice struct {
	db *gorm.DB
}

func NewFilingCabinetService() FilingCabinetService {
	return &filingcabinetservice{
		db: config.DB,
	}
}

func (s *filingcabinetservice) GetFilingCabinet(ctx context.Context, cabinetID int) ([]model.FilingCabinet, error) {
	if cabinetID <= 0 {
		return nil, apperror.Invalid("cabinet id bust be a positive integer", nil)
	}

	ctx, cancel := context.WithTimeout(ctx, utils.DefaultQueryTimeout)
	defer cancel()

	var filingcabinet []model.FilingCabinet
	err := s.db.WithContext(ctx).Where("cabinet_id = ?", cabinetID).Find(&filingcabinet).Error
	if err != nil {
		return nil, apperror.Internal("failed to fetch filingcabinet", err)
	}
	return filingcabinet, nil
}
