package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Import the database driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Import the file source driver
	"go-to-chat/app/config"
	"log"
)

func MigrateDB(dbConfig *config.DatabaseConfig) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	migrations, err := migrate.New("file://database/migrations", dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	err = migrations.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	if err == nil {
		log.Println("Database migration successful")
	}
}
