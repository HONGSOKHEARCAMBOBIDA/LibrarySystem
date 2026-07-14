package model

import "mysql/model/base"

type Author struct {
	base.ModelBase
	Name     string `json:"name" gorm:"column:name"`
	Gender   int    `json:"gender" gorm:"column:gender"`
	Isactive bool   `json:"is_active" gorm:"column:is_active"`
	Note     string `json:"note" gorm:"column:note"`
}

func (Author) TableName() string {
	return "authors"
}
