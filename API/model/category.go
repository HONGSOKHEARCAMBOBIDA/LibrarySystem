package model

import "mysql/model/base"

type Category struct {
	base.ModelBase
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Isactive    bool   `json:"is_active" gorm:"column:is_active"`
}

func (Category) TableName() string {
	return "category"
}
