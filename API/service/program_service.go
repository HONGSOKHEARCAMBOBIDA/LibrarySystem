package service

import (
	"context"
	"mysql/config"
	"mysql/constant/apperror"
	"mysql/model"
	"time"

	"gorm.io/gorm"
)

type ProgramService interface {
	GetProgram(ctx context.Context, departmentID int) ([]model.Program, error)
}

type programservice struct {
	db *gorm.DB
}

func NewProgramService() ProgramService {
	return &programservice{
		db: config.DB,
	}
}

func (s *programservice) GetProgram(ctx context.Context, departmentID int) ([]model.Program, error) {
	if departmentID <= 0 {
		return nil, apperror.Invalid("department id must be a positive integer", nil)
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var programs []model.Program
	err := s.db.WithContext(ctx).Where("department_id = ?", departmentID).Find(&programs).Error
	if err != nil {
		return nil, apperror.Internal("failed to fetch program", err)
	}
	return programs, nil
}
