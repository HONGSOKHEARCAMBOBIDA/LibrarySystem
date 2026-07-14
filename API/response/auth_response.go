package response

import (
	"mysql/model"
	"mysql/model/base"
)

type AuthResponse struct {
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	AccessToken  string             `json:"token"`
	RefreshToken string             `json:"refreshToken"`
	Role         []model.Role       `json:"roles"`
	Permissions  []model.Permission `json:"permissions"`
}

type UserDataResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	RoleID      int                `json:"role_id" gorm:"column:role_id"`
	Permissions []model.Permission `json:"permissions" gorm:"-"`
}

type UserResponse struct {
	base.ModelBase
	NameKh   string `json:"name_kh" gorm:"column:name_kh"`
	NameEN   string `json:"name_en" gorm:"column:name_en"`
	Email    string `json:"email" gorm:"column:email"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	Gender   int    `json:"gender" gorm:"column:gender"`
	Dob      string `json:"dob" gorm:"column:dob"`
	Isactive bool   `json:"is_active" gorm:"column:is_active"`
}
