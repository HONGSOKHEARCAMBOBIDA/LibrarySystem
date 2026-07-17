package seed

import (
	"mysql/constant/seeddata"
	"mysql/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedCategory(db *gorm.DB) error {

	var category []model.Category

	for _, c := range seeddata.Categories {
		category = append(category, model.Category{
			Name:        c.Name,
			Description: c.Description,
			Isactive:    c.Isactive,
		})
	}

	return db.Clauses(clause.OnConflict{
		// OnConflict is Upsert = Update + Insert
		Columns: []clause.Column{
			{Name: "name"}, // unique key
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"description",
		}),
	}).Create(&category).Error
}
