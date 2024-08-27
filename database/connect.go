package database

import (
	"fmt"
	"go-to-chat/app/config"
	"go-to-chat/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dbInstance *gorm.DB

func SetupDatabase() {
	appConfig, err := config.GetAppConfig()
	if err != nil {
		log.Fatal("Error getting app config: ", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		appConfig.Database.Host,
		appConfig.Database.Username,
		appConfig.Database.Password,
		appConfig.Database.DBName,
		appConfig.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = migration(db)
	if err != nil {
		panic("failed to migrate database")
	}

	log.Println("Database connected")

	dbInstance = db
}

func GetDBInstance() *gorm.DB {
	return dbInstance

}

func migration(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		return err
	}
	return nil
}
