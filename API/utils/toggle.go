package utils

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func ToggleStatus[T any](ctx context.Context, db *gorm.DB, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid id")
	}

	result := db.WithContext(ctx).Model(new(T)).
		Where("id = ?", id).
		Update("is_active", gorm.Expr("NOT is_active"))
	return result.Error
}
