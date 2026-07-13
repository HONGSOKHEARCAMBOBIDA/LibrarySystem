package helper

import (
	"mysql/model"

	"gorm.io/gorm"
)

func ApplyAccessGetRole(query, db *gorm.DB, role model.Role, user model.User) *gorm.DB {
	return query.Where("r.level <= ?", user.Role.Level)
}

func ApplyAccessGetUser(query, db *gorm.DB, user model.User) *gorm.DB {
	return query.Where("u.level <= ?", user.Level)
}
