package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	username := os.Getenv("DB_USERNAME")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASS")
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
