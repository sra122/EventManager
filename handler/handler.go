package handler

import (
	"example.com/hello/dbconnection"
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

func initialise() Handler {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectDb()
	h := New(DB)

	return h
}
