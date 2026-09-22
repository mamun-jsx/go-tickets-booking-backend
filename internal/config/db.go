package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDatabase(cfg *Config) *gorm.DB {
	dsn := cfg.Dns
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("Failed to conncet databae")
	} else {
		println("===================================")
		println("______Database connected______")
		println("===================================")
	}
	return db
}
