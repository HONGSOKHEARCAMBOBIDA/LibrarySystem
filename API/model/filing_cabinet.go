package model

import "mysql/model/base"

type FilingCabinet struct {
	base.ModelBase
	Name      string `json:"name" gorm:"column:name"`
	Isactive  bool   `json:"is_active" gorm:"column:is_active"`
	CabinetID int    `json:"cabinet_id" gorm:"column:cabinet_id"`
}

func (FilingCabinet) TableName() string {
	return "filing_cabinet"
}
