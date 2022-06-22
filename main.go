package main

import (
	"example.com/hello/dbconnection"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	Conn   *gorm.DB
	Router *mux.Router
	Logger *log.Logger
}

func (app *App) initializeRouter() {
	r := app.Router

	//Employee
	r.HandleFunc("/employees", app.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", app.GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{employee_id}", app.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{employee_id}", app.DeleteEmployee).Methods("DELETE")

	//Event
	r.HandleFunc("/events", app.CreateEvent).Methods("POST")
	r.HandleFunc("/event/{event_id}", app.GetEvent).Methods("GET")

	//EmployeeEvent
	r.HandleFunc("/event/{event_id}/employees", app.AddEmployeeForEvent).Methods("POST")
	r.HandleFunc("/event/{event_id}/employees", app.GetEmployeesForEvent).Methods("GET")
	r.HandleFunc("/event/{event_id}/employees", app.GetEmployeesForEvent).Methods("GET").Queries("is_accommodation_required", "{accommodation}")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func (app *App) InitialMigration() {

	DB, error := dbconnection.ConnectDb()

	if error == nil {
		app.Conn = DB
		DB.AutoMigrate(&Employee{})
		DB.AutoMigrate(&Event{})
	}
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	a := App{}
	a.InitialMigration()
	a.Router = mux.NewRouter()
	a.initializeRouter()
}
