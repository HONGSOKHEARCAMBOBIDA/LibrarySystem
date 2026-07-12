package response

import "mysql/model"

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
