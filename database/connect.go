package database

import (
	"fmt"
	"go-to-chat/app/config"
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

	db, err := CreateDataBase(appConfig.Database)
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Database connected")

	MigrateDB(appConfig.Database)
	dbInstance = db
}

func GetDBInstance() *gorm.DB {
	return dbInstance
}

func CreateDataBase(dbConfig *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
