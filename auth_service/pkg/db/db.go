package db

import (
	"Auth/pkg/config"
	"Auth/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.Employer{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.JobSeeker{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.Connections{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.JobseekerSubscription{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.JobseekerPlan{}); err != nil {
		return nil, err
	}

	return db, nil
}