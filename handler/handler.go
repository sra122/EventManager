package handler

import (
	"example.com/hello/dbconnection"
	"example.com/hello/model"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

// Initialise for UnitTests
func Initialise() Handler {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectTestDb()
	h := New(DB)

	return h
}

// Drop the table after testing.
func DropTable(handler Handler) {

	handler.DB.Migrator().DropTable(&model.Employee{})
	handler.DB.Migrator().DropTable(&model.Event{})
	handler.DB.Migrator().DropTable(&model.EventEmployees{})
}
