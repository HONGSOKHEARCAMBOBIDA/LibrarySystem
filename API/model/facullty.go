package model

import "mysql/model/base"

type Faculty struct {
	base.ModelBase
	Name     string `json:"name" gorm:"column:name"`
	Code     string `json:"code" gorm:"column:code;uniqueIndex"`
	Isactive bool   `json:"is_active" gorm:"column:is_active"`
}

func (Faculty) TableName() string {
	return "faculties"
}
