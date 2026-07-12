package model

import "mysql/model/base"

type User struct {
	base.ModelBase
	NameKh   string `json:"name_kh" gorm:"column:name_kh"`
	NameEN   string `json:"name_en" gorm:"column:name_en"`
	Email    string `json:"email" gorm:"column:email;uniqueIndex"`
	Password string `json:"password" gorm:"column:password"`
	Level    int    `json:"level" gorm:"column:level"`
	RoleId   int    `json:"role_id"`
	Gender   int    `json:"gender" gorm:"column:gender"`
	Dob      string `json:"dob" gorm:"column:dob"`
	Isactive bool   `json:"is_active" gorm:"column:is_active"`
	Role     Role
}

func (User) TableName() string {
	return "user"
}
