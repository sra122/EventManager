package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type birthDay time.Time

type Employee struct {
	gorm.Model        //TODO: Add email for converting unique record.
	FirstName               string   `json:"firstName"`
	LastName string   `json:"lastName"`
	BirthDay birthDay `json:"birthDay"`
	Gender   string   `json:"gender"` // Gender
	IsAccommodationRequired bool     `json:"is_accommodation_required"`
}

// GetEmployees
// Get List of Employees in the Organization/**
func (app App) GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee []Employee
	app.Conn.Order("id asc").Find(&employee)
	json.NewEncoder(w).Encode(employee)
}

// CreateEmployee
// Create an Employee to the Organization
func (app App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	requestBodyError := json.NewDecoder(r.Body).Decode(&employee)
	if requestBodyError != nil {
		// If request body doesn't fit according to the requirements
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(strings.Trim(requestBodyError.Error(), "\""))
		return
	}
	error := app.Conn.Create(&employee).Error
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(employee)
	} else {
		//Error occurred during creating of employee
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.Error())
	}
}

// UpdateEmployee
// Updates the Employee for the provided id in the url
func (app App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	notFoundError := app.Conn.First(&employee, params["employee_id"]).Error
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Employee record is not found with id " + params["employee_id"])
		return
	}
	requestBodyError := json.NewDecoder(r.Body).Decode(&employee)
	if requestBodyError != nil {
		// If requestbody doesn't fit according to the requirements
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(strings.Trim(requestBodyError.Error(), "\""))
		return
	}
	error := app.Conn.Save(&employee).Error
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(employee)
	} else {
		// Error occurred during the save of Entity
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(error.Error())
	}
}

// DeleteEmplpoyee
// Delete Employee from the records
func (app App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	notFoundError := app.Conn.First(&employee, params["employee_id"])
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Employee record is not found with id " + params["employee_id"])
		return
	}

	error := app.Conn.Delete(&employee, params["employee_id"]).Error
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Employee is deleted ||")
	} else {
		// Error occurred during the delete operation.
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(error.Error())
	}

}

func (j *birthDay) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = birthDay(t)
	return nil
}

func (j birthDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
