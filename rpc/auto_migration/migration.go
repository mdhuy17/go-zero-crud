package migration

import (
	"github.com/pkg/errors"
	"go_zero-crud/rpc/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Company{},
	); err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}

	return nil
}
