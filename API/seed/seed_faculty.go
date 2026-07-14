package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedFacultiess(db *gorm.DB) error {

	var faculties []model.Faculty

	for _, f := range seeddata.Faculties {
		faculties = append(faculties, model.Faculty{
			Name:     f.Name,
			Code:     f.Code,
			Isactive: true,
		})
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "code"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"is_active",
		}),
	}).Create(&faculties).Error
}
