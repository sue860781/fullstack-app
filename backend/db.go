package db

import {
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
}

func Connect() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=appdb port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
