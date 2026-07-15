package service

import (
	"context"
	"errors"
	"mysql/config"
	"mysql/model"
	"mysql/request"
	"mysql/response"

	"gorm.io/gorm"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, input request.AuthorRequestCreate) error
	GetAuthor(ctx context.Context, pf request.Pagination, filter map[string]string) ([]response.AuthorResponse, *model.PaginationMetadata, error)
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
