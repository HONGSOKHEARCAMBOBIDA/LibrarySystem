package model

import "mysql/model/base"

type Cabinet struct {
	base.ModelBase
	Name     string `json:"name" gorm:"column:name"`
	Isactive bool   `json:"is_active" gorm:"column:is_active"`
}

func (Cabinet) TableName() string {
	return "cabinet"
}
