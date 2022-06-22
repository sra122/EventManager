package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EventEmployees struct {
	EmployeeId int `json:"employee_id"`
	EventId    int `json:"event_id"`
}

// AddEmployeeForEvent
// Add an employee to the event.
func (app App) AddEmployeeForEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get the params from the url
	var eventEmployees EventEmployees
	//Convert String value to integer
	eventId, conversionError := strconv.Atoi(params["event_id"])
	if conversionError != nil {
		// Error during the conversion
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(conversionError.Error())
		return
	}
	eventEmployees.EventId = eventId

	requestBodyError := json.NewDecoder(r.Body).Decode(&eventEmployees)
	if requestBodyError != nil {
		// Error in the request body
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestBodyError.Error())
		return
	}

	err := app.Conn.Create(&eventEmployees).Error
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(eventEmployees)
	} else {
		// Error while creating the entity
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(err.Error())
	}
}

// GetEmployeesForEvent
// Get the list of Employees for a particular Event.
func (app App) GetEmployeesForEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	accommodationQuery := r.URL.Query().Get("is_accommodation_required")
	var event Event
	var err error
	notFoundError := app.Conn.First(&event, params["event_id"]).Error
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Event record is not found with id " + params["event_id"])
		return
	}

	if accommodationQuery != "" {
		err = app.Conn.Preload("Employees", "is_accommodation_required", accommodationQuery).Find(&event, params["event_id"]).Error
	} else {
		err = app.Conn.Preload("Employees").Find(&event, params["event_id"]).Error
	}

	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(event)
	} else {
		// Error while fetching the records
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
}
