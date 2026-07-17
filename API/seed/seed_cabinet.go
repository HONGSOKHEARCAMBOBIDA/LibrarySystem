package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCabinet(db *gorm.DB) error {

	var cabinets []model.Cabinet

	for _, c := range seeddata.Cabinets {
		cabinets = append(cabinets, model.Cabinet{
			Name:     c.Name,
			Isactive: c.Isactive,
		})
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "name"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
		}),
	}).Create(&cabinets).Error
}
