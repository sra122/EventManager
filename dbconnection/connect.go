package dbconnection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDb() *gorm.DB {

	var DB *gorm.DB
	var err error

	host := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, userName, password, dbName, dbPort)
	//var DNS = "postgres://postgres:postgres@db:5432/postgres"
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	return DB
}

func closeDb() {

}
