package seed

import (
	"mysql/constant/seeddata"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedUser(db *gorm.DB) error {
	user := seeddata.DeveloperUser

	return db.
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "email"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"name_kh",
				"name_en",
				"role_id",
				"level",
				"gender",
				"dob",
				"is_active",
			}),
		}).
		Create(&user).
		Error
}
