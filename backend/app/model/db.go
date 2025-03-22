package model

import (
	"fmt"

	"github.com/dawnseeker8/questionnaire-system/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	// Auto migrate models
	err = DB.AutoMigrate(
		&Form{},
		&Section{},
		&Question{},
		&QuestionOption{},
		&Response{},
		&Answer{},
	)
	if err != nil {
		return err
	}

	return nil
}
