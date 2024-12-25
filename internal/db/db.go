package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"santa25-52/internal/config"
)

type Member struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Wish string
	TgID int64
}

func MustLoad(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&Member{})
	if err != nil {
		panic("Failed to auto-migrate")
	}

	return db
}
