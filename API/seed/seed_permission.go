package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedPermissions(db *gorm.DB) error {

	var permissions []model.Permission

	for _, p := range seeddata.Permissions {
		permissions = append(permissions, model.Permission{
			Name:        p.Name,
			ModuleName:  p.ModuleName,
			DisplayName: p.DisplayName,
			OrderNo:     p.OrderNo,
		})
	}

	return db.Clauses(clause.OnConflict{
		// OnConflict is Upsert = Update + Insert
		Columns: []clause.Column{
			{Name: "name"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"module_name",
			"display_name",
			"order_no",
		}),
	}).Create(&permissions).Error
}
