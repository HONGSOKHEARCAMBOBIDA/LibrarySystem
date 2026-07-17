package service

import (
	"context"
	"errors"
	"mysql/config"
	"mysql/constant/apperror"
	"mysql/model"
	"mysql/request"
	"mysql/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CategoryService interface {
	GetCategory(ctx context.Context) ([]model.Category, error)
	CreateCategory(ctx context.Context, input request.CategoryRequestCreate) error
	UpdateCategory(ctx context.Context, id int, input request.CategoryRequestUpdate) error
	ToggleStatusCategory(ctx context.Context, id int) error
}

const defaultQueryTimeout = 5 * time.Second

type categoryservice struct {
	db *gorm.DB
}

func NewCategoryService() CategoryService {
	return &categoryservice{
		db: config.DB,
	}
}

func (s *categoryservice) GetCategory(ctx context.Context) ([]model.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var category []model.Category
	err := s.db.WithContext(ctx).Find(&category).Error
	if err != nil {
		return nil, apperror.Internal("failed to fetch categorys", err)
	}
	return category, nil
}

func (s *categoryservice) CreateCategory(ctx context.Context, input request.CategoryRequestCreate) error {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return apperror.Invalid("category name is required", nil)
	}

	ctx, cancel := context.WithTimeout(ctx, defaultQueryTimeout)
	defer cancel()

	newcategory := model.Category{
		Name:        name,
		Description: input.Description,
		Isactive:    true,
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&newcategory).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return apperror.New(apperror.CodeConflict, "category name already exists", err)
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return apperror.New(apperror.CodeUnavailable, "request timed out, please try again", err)
		}
		return apperror.Internal("failed to create category", err)
	}
	return nil
}

func (s *categoryservice) UpdateCategory(ctx context.Context, id int, input request.CategoryRequestUpdate) error {
	if id <= 0 {
		return apperror.Invalid("category id must be a positive integer", nil)
	}

	updates := map[string]interface{}{}

	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		if name == "" {
			return apperror.Invalid("category name cannot be empty", nil)
		}
		updates["name"] = name
	}
	if input.Description != nil {
		updates["description"] = *input.Description
	}

	if len(updates) == 0 {
		return apperror.Invalid("no fields to update", nil)
	}

	ctx, cancel := context.WithTimeout(ctx, defaultQueryTimeout)
	defer cancel()

	var rowsAffected int64
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Category{}).
			Where("id = ?", id).
			Updates(updates)
		rowsAffected = result.RowsAffected
		return result.Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return apperror.New(apperror.CodeConflict, "category name already exists", err)
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return apperror.New(apperror.CodeUnavailable, "request timed out, please try again", err)
		}
		return apperror.Internal("failed to update category", err)
	}

	if rowsAffected == 0 {
		return apperror.NotFound("category not found", nil)
	}

	return nil
}

func (s *categoryservice) ToggleStatusCategory(ctx context.Context, id int) error {
	return utils.ToggleStatus[model.Category](ctx, s.db, id)
}
