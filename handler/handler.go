package handler

import (
	"example.com/hello/dbconnection"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func initialise() handler {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectDb()
	h := New(DB)

	return h
}
