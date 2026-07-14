package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedPrograms(db *gorm.DB) error {

	var programs []model.Program

	for _, p := range seeddata.Programs {
		programs = append(programs, model.Program{
			Name:         p.Name,
			Code:         p.Code,
			Isactive:     p.Isactive,
			DepartmentID: p.DepartmentID,
		})
	}

	return db.Clauses(clause.OnConflict{
		// OnConflict is Upsert = Update + Insert
		Columns: []clause.Column{
			{Name: "code"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"department_id",
		}),
	}).Create(&programs).Error
}
