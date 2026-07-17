package service

import (
	"context"
	"mysql/config"
	"mysql/constant/apperror"
	"mysql/model"
	"time"

	"gorm.io/gorm"
)

type DepartmentService interface {
	GetDepartment(ctx context.Context, facultyID int) ([]model.Department, error)
}

type departmentservice struct {
	db *gorm.DB
}

func NewDepartmentService() DepartmentService {
	return &departmentservice{
		db: config.DB,
	}
}

func (s *departmentservice) GetDepartment(ctx context.Context, facultyID int) ([]model.Department, error) {
	if facultyID <= 0 {
		return nil, apperror.Invalid("faculty id must be a positive integer", nil)
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var departments []model.Department
	err := s.db.WithContext(ctx).
		Where("faculty_id = ?", facultyID).
		Find(&departments).Error
	if err != nil {
		return nil, apperror.Internal("failed to fetch departments", err)
	}
	return departments, nil
}
