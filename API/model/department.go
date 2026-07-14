package model

import "mysql/model/base"

type Department struct {
	base.ModelBase
	Name      string `json:"name" gorm:"column:name"`
	Code      string `json:"code" gorm:"column:code;uniqueIndex"`
	Isactive  bool   `json:"is_active" gorm:"column:is_active"`
	FacultyID int    `json:"faculty_id" gorm:"column:faculty_id"`
}

func (Department) TableName() string {
	return "department"
}
