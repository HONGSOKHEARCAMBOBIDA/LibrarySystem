package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedFilingCabinet(db *gorm.DB) error {

	var filingcabinet []model.FilingCabinet

	for _, p := range seeddata.FilingCabinets {
		filingcabinet = append(filingcabinet, model.FilingCabinet{
			Name:      p.Name,
			Isactive:  p.Isactive,
			CabinetID: p.CabinetID,
		})
	}

	return db.Clauses(clause.OnConflict{
		// OnConflict is Upsert = Update + Insert
		Columns: []clause.Column{
			{Name: "name"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
		}),
	}).Create(&filingcabinet).Error
}
