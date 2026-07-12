package model

type Permission struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100;not null;uniqueIndex" json:"name"`
	ModuleName  string `json:"module_name" gorm:"column:module_name"`
	DisplayName string `json:"display_name" gorm:"column:display_name"`
	OrderNo     int    `json:"order_no" gorm:"column:order_no"`
}

func (Permission) TableName() string {
	return "permission"
}
