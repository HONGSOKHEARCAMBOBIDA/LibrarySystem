package service

import (
	"context"
	"fmt"
	"mysql/config"
	"mysql/model"
	"mysql/request"
	"mysql/response"
	"mysql/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, input request.AuthorRequestCreate) error
	UpdateAuthor(ctx context.Context, id int, input request.AuthorRequestUpdate) error
	GetAuthor(ctx context.Context, pf request.Pagination, filter map[string]string) ([]response.AuthorResponse, *model.PaginationMetadata, error)
	ToggleStatusAuthor(ctx context.Context, id int) error
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
				FacultyID: *input.FacultyIDs[i],
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

func (s *authorservice) GetAuthor(ctx context.Context, pf request.Pagination, filter map[string]string) ([]response.AuthorResponse, *model.PaginationMetadata, error) {
	var authors []response.AuthorResponse
	offset := (pf.Page - 1) * pf.PageSize

	applyFilters := func(tx *gorm.DB) *gorm.DB {
		if v, ok := filter["name"]; ok && v != "" {
			tx = tx.Where("a.name LIKE ?", "%"+v+"%")
		}
		if v, ok := filter["gender"]; ok && v != "" {
			tx = tx.Where("a.gender = ?", v)
		}
		if v, ok := filter["is_active"]; ok && v != "" {
			tx = tx.Where("a.is_active = ?", v)
		}
		return tx
	}

	var total int64
	countQuery := applyFilters(s.db.WithContext(ctx).Table("authors a"))
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	dataQuery := applyFilters(s.db.WithContext(ctx).Table("authors a")).
		Select(`
			a.id AS id,
			a.name AS name,
			a.gender AS gender,
			a.is_active AS is_active,
			a.note AS note
		`)

	if err := dataQuery.Offset(offset).Limit(pf.PageSize).Scan(&authors).Error; err != nil {
		return nil, nil, err
	}

	metadata := &model.PaginationMetadata{
		Page:       pf.Page,
		PageSize:   pf.PageSize,
		TotalCount: total,
	}

	if len(authors) == 0 {
		return authors, metadata, nil
	}

	authorIDs := make([]int, len(authors))
	for i, a := range authors {
		authorIDs[i] = a.ID
	}

	type facultyRow struct {
		AuthorID int
		ID       int
		Name     string
		Code     string
	}

	var facultyRows []facultyRow
	if err := s.db.WithContext(ctx).Table("author_faculty af").
		Select(`
			af.author_id AS author_id,
			f.id AS id,
			f.name AS name,
			f.code AS code
		`).
		Joins("JOIN faculties f ON f.id = af.faculty_id").
		Where("af.author_id IN ?", authorIDs).
		Scan(&facultyRows).Error; err != nil {
		return nil, nil, err
	}

	facultyMap := make(map[int][]model.Faculty, len(authors))
	for _, fr := range facultyRows {
		facultyMap[fr.AuthorID] = append(facultyMap[fr.AuthorID], model.Faculty{
			ID:   fr.ID,
			Name: fr.Name,
			Code: fr.Code,
		})
	}

	for i := range authors {
		authors[i].Faculty = facultyMap[authors[i].ID]
	}

	return authors, metadata, nil
}

func (s *authorservice) UpdateAuthor(ctx context.Context, id int, input request.AuthorRequestUpdate) error {
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	committed := false
	defer func() {
		if r := recover(); r != nil || !committed {
			tx.Rollback()
		}
	}()

	var author model.Author
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&author, id).Error; err != nil {
		return err
	}

	updates := map[string]interface{}{}
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.Gender != nil {
		updates["gender"] = *input.Gender
	}
	if input.Note != nil {
		updates["note"] = *input.Note
	}

	if len(updates) > 0 {
		if err := tx.Model(&model.Author{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return fmt.Errorf("failed to update author: %w", err)
		}
	}

	if len(input.FacultyIDs) != 0 {
		if err := tx.Where("author_id = ?", author.ID).Delete(&model.AuthorFaculty{}).Error; err != nil {
			return fmt.Errorf("failed to clear old faculties: %w", err)
		}

		authorfaculties := make([]model.AuthorFaculty, 0, len(input.FacultyIDs))
		for _, fid := range input.FacultyIDs {
			if fid == nil {
				continue
			}
			authorfaculties = append(authorfaculties, model.AuthorFaculty{
				AuthorID:  author.ID,
				FacultyID: *fid,
			})
		}

		if len(authorfaculties) > 0 {
			if err := tx.Create(&authorfaculties).Error; err != nil {
				return fmt.Errorf("failed to create author faculties: %w", err)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true
	return nil
}

func (s *authorservice) ToggleStatusAuthor(ctx context.Context, id int) error {
	return utils.ToggleStatus[model.Author](ctx, s.db, id)
}
