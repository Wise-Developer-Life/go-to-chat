package database

import (
	"go-to-chat/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dbInstance *gorm.DB

func SetupDatabase() {
	dsn := "host=localhost user=postgres password=secret dbname=postgres port=5434"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = migration(db)
	if err != nil {
		panic("failed to migrate database")
	}

	// log.Println("Database connected")
	log.Println("Database connected")

	dbInstance = db
}

func GetDBInstance() *gorm.DB {
	return dbInstance

}
func isDbAvailable() bool {
	return dbInstance != nil
}

func migration(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		return err
	}
	return nil
}
