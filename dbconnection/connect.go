package dbconnection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDb() (*gorm.DB, error) {

	var DB *gorm.DB
	var err error

	host := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var DNS = "host=" + host + " user=" + userName + " password=" + password + " dbname=" + dbName + " port=" + dbPort
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	return DB, err
}

func closeDb() {

}
