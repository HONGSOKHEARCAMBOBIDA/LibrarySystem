package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedRoles(db *gorm.DB) error {

	var roles []model.Role

	for _, p := range seeddata.Roles {
		roles = append(roles, model.Role{
			Name:       p.Name,
			ModuleName: p.ModuleName,
			IsActive:   true,
			Level:      p.Level,
		})
	}

	return db.Clauses(clause.OnConflict{
		// OnConflict is Upsert = Update + Insert
		Columns: []clause.Column{
			{Name: "name"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"module_name",
			"level",
		}),
	}).Create(&roles).Error
}
