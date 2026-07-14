package request

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	NameKh string `json:"name_kh"`
	NameEN string `json:"name_en"`
	RoleID int    `json:"role_id"`
	Gender int    `json:"gender"`
	Dob    string `json:"dob"`
}

type UpdateUserRequest struct {
	NameKh *string `json:"name_kh"`
	NameEN *string `json:"name_en"`
	RoleID *int    `json:"role_id"`
	Gender *int    `json:"gender"`
	Dob    *string `json:"dob"`
}
