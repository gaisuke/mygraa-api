package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitializeDB() {
	dsn := "host=localhost user=postgres password=root dbname=mygram_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// for initial run, migrating the database
	// db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
