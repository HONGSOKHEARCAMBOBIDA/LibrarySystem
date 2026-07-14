package model

import "mysql/model/base"

type AuthorFaculty struct {
	base.ModelBase
	FacultyID int `json:"faculty_id" gorm:"column:faculty_id"`
	AuthorID  int `json:"author_id" gorm:"column:author_id"`
}

func (AuthorFaculty) TableName() string {
	return "author_faculty"
}
