package test

import (
	"example.com/hello/dbconnection"
	"example.com/hello/pkg/employee"
	"example.com/hello/pkg/employee_event"
	"example.com/hello/pkg/event"
	"github.com/joho/godotenv"
	"log"
)

func InitializeEmployeeDBConnection() (*employee.Task, *employee.EmployeeConnection) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectTestDb()

	empRepo := employee.InitialiseEmployeeHandler(DB)
	empTask := employee.NewTask(empRepo)

	return empTask, empRepo
}

func DropEmployeeTable(conn employee.EmployeeConnection) {
	err := conn.DB.Migrator().DropTable(&employee.Employee{})
	if err != nil {
		return
	}
}

func InitializeEventDBConnection() (*event.Task, *event.EventConnection) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectTestDb()

	eveRepo := event.InitialiseEventHandler(DB)
	eveTask := event.NewTask(eveRepo)

	return eveTask, eveRepo
}

func DropEventTable(conn event.EventConnection) {
	err := conn.DB.Migrator().DropTable(&event.Event{})
	if err != nil {
		return
	}
}

func InitializeEmployeeEventDBConnection() (*employee_event.Task, *employee_event.EmployeeEventConnection) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectTestDb()

	empEveRepo := employee_event.InitialiseEmployeeEventHandler(DB)
	empEveTask := employee_event.NewTask(empEveRepo)

	return empEveTask, empEveRepo
}

func DropEmployeeEventTable(conn employee_event.EmployeeEventConnection) {
	err := conn.DB.Migrator().DropTable(&employee_event.EventEmployees{})
	if err != nil {
		return
	}
}
