package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedDepartments(db *gorm.DB) error {

	var departments []model.Department

	for _, d := range seeddata.Departments {
		departments = append(departments, model.Department{
			Name:      d.Name,
			Code:      d.Code,
			Isactive:  true,
			FacultyID: d.FacultyID,
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
	}).Create(&departments).Error
}
