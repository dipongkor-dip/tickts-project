package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *Config) *gorm.DB {
	var db, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		panic("❌ failed to connect database")
	} else {
		fmt.Println("✅ Database connected successfully")
	}

	return db
}
