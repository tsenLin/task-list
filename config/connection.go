package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	// Connect to an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(Conf.GetString("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}