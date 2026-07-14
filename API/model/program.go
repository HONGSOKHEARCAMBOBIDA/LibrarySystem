package model

import "mysql/model/base"

type Program struct {
	base.ModelBase
	Name         string `json:"name" gorm:"column:name"`
	Code         string `json:"code" gorm:"column:code;uniqueIndex"`
	Isactive     bool   `json:"is_active" gorm:"column:is_active"`
	DepartmentID int    `json:"department_id" gorm:"column:department_id"`
}

func (Program) TableName() string {
	return "program"
}
