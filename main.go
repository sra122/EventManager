package main

import (
	"example.com/hello/dbconnection"
	"example.com/hello/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//Load Environment Variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	initialiseMain()
}

func initialiseMain() {
	DB := dbconnection.ConnectDb()
	h := handler.New(DB)

	router := mux.NewRouter()

	//Employee
	router.HandleFunc("/employees", h.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees", h.GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{employee_id}", h.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{employee_id}", h.DeleteEmployee).Methods("DELETE")

	//Event
	router.HandleFunc("/events", h.CreateEvent).Methods("POST")
	router.HandleFunc("/events", h.GetUpcomingEvents).Methods("GET")
	router.HandleFunc("/event/{event_id}", h.GetEvent).Methods("GET")

	//EmployeeEvent
	router.HandleFunc("/event/{event_id}/employees", h.AddEmployeeForEvent).Methods("POST")
	router.HandleFunc("/event/{event_id}/employees", h.GetEmployeesForEvent).Methods("GET")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
