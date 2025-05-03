package database

import (
	"fmt"

	"github.com/bit365/iotcloud/services/webapi/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, err
}
