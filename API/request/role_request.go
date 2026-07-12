package request

type RoleRequestUpdate struct {
	Name       *string `json:"name"`
	ModuleName *string `json:"module_name" gorm:"column:module_name"`
}
