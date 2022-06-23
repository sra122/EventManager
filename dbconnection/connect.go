package dbconnection

import (
	"example.com/hello/model"
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
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	DB.AutoMigrate(&model.Employee{})
	DB.AutoMigrate(&model.Event{})
	DB.AutoMigrate(&model.EventEmployees{})

	return DB
}

func ConnectTestDb() *gorm.DB {

	var DB *gorm.DB
	var err error

	host := os.Getenv("TEST_DB_HOST")
	userName := os.Getenv("TEST_DB_USERNAME")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")
	dbPort := os.Getenv("TEST_DB_PORT")

	var DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, userName, password, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	DB.Migrator().DropTable(&model.Employee{})
	DB.Migrator().DropTable(&model.Event{})
	DB.Migrator().DropTable(&model.EventEmployees{})

	DB.AutoMigrate(&model.Employee{})
	DB.AutoMigrate(&model.Event{})
	DB.AutoMigrate(&model.EventEmployees{})

	return DB
}
