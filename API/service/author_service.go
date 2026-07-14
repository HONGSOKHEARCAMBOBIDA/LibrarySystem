package service

import (
	"context"
	"errors"
	"mysql/config"
	"mysql/model"
	"mysql/request"

	"gorm.io/gorm"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, input request.AuthorRequestCreate) error
}

type authorservice struct {
	db *gorm.DB
}

func NewAuthorService() AuthorService {
	return &authorservice{
		db: config.DB,
	}
}

func (s *authorservice) CreateAuthor(ctx context.Context, input request.AuthorRequestCreate) error {
	if len(input.FacultyIDs) == 0 {
		return errors.New("faculty cannot be empty")
	}
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	committed := false
	defer func() {
		if r := recover(); r != nil || !committed {
			tx.Rollback()
		}
	}()

	newauthor := model.Author{
		Name:     input.Name,
		Gender:   input.Gender,
		Isactive: true,
		Note:     input.Note,
	}
	if err := tx.Create(&newauthor).Error; err != nil {
		return err
	}

	if len(input.FacultyIDs) != 0 {
		for i := range input.FacultyIDs {
			authorfaculty := model.AuthorFaculty{
				FacultyID: input.FacultyIDs[i],
				AuthorID:  newauthor.ID,
			}
			if err := tx.Create(&authorfaculty).Error; err != nil {
				return err
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true
	return nil
}
