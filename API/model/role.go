package model

import "mysql/model/base"

type Role struct {
	base.ModelBase
	Name        string       `json:"name" gorm:"uniqueIndex"`
	ModuleName  string       `json:"module_name" gorm:"column:module_name"`
	IsActive    bool         `json:"is_active" gorm:"column:is_active"`
	Permissions []Permission `gorm:"many2many:role_has_permissions"`
	Level       int          `json:"level" gorm:"column:level"`
}

func (Role) TableName() string {
	return "role"
}
