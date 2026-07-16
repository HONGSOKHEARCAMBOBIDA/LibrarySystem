package service

import (
	"context"

	"mysql/config"
	"mysql/model"

	"gorm.io/gorm"
)

type FacultyService interface {
	GetFaculty(ctx context.Context) ([]model.Faculty, error)
}

type facultyservice struct {
	db *gorm.DB
}

func NewFacultyService() FacultyService {
	return &facultyservice{
		db: config.DB,
	}
}

func (s *facultyservice) GetFaculty(ctx context.Context) ([]model.Faculty, error) {
	var faculties []model.Faculty

	if err := s.db.WithContext(ctx).Find(&faculties).Error; err != nil {
		return nil, err
	}

	return faculties, nil
}
