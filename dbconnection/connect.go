package dbconnection

import (
	"example.com/hello/pkg/employee"
	"example.com/hello/pkg/employee_event"
	"example.com/hello/pkg/event"
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

	DB.AutoMigrate(&employee.Employee{})
	DB.AutoMigrate(&event.Event{})
	DB.AutoMigrate(&employee_event.EventEmployees{})

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

	DB.Migrator().DropTable(&employee.Employee{})
	DB.Migrator().DropTable(&event.Event{})
	DB.Migrator().DropTable(&employee_event.EventEmployees{})

	DB.AutoMigrate(&employee.Employee{})
	DB.AutoMigrate(&event.Event{})
	DB.AutoMigrate(&employee_event.EventEmployees{})

	return DB
}
