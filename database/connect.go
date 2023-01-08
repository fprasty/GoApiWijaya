package database

import (
	"log"
	"os"

	"github.com/fprasty/GoApiWijaya/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Can't load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}
	DB = database

	database.AutoMigrate(
		&models.User{},
		//&models.UserBarang{},
		//&models.BarangComment{},
		&models.Admin{},
	)
}
