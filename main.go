package main

import (
	"example.com/hello/dbconnection"
	"example.com/hello/pkg/employee"
	"example.com/hello/pkg/employee_event"
	"example.com/hello/pkg/event"
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
	router := mux.NewRouter()

	//Employee
	empRepo := employee.InitialiseEmployeeHandler(DB)
	empTask := employee.NewTask(empRepo)

	router.HandleFunc("/employees", empTask.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees", empTask.GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{employee_id}", empTask.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{employee_id}", empTask.DeleteEmployee).Methods("DELETE")

	//Event
	eventRepo := event.InitialiseEventHandler(DB)
	eventTask := event.NewTask(eventRepo)

	router.HandleFunc("/events", eventTask.CreateEvent).Methods("POST")
	router.HandleFunc("/events", eventTask.GetUpcomingEvents).Methods("GET")
	router.HandleFunc("/event/{event_id}", eventTask.GetEvent).Methods("GET")

	//EmployeeEvent
	empEventRepo := employee_event.InitialiseEmployeeEventHandler(DB)
	empEventTask := employee_event.NewTask(empEventRepo)

	router.HandleFunc("/event/{event_id}/employees", empEventTask.AddEmployeeForEvent).Methods("POST")
	router.HandleFunc("/event/{event_id}/employees", empEventTask.GetEmployeesForEvent).Methods("GET")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
