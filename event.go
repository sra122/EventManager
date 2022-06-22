package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type eventDate time.Time

type Event struct {
	gorm.Model
	Name      string     `json:"name"`
	Date      eventDate  `json:"date"`
	Employees []Employee `gorm:"many2many:event_employees;"`
}

// CreateEvent
func (app App) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Event
	requestBodyError := json.NewDecoder(r.Body).Decode(&event)
	if requestBodyError != nil {
		// Error in the requestbody.
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(requestBodyError.Error())
		return
	}

	error := app.Conn.Create(&event).Error
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(event)
	} else {
		// Error while creating an entity
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.Error())
	}
}

// GetEvent
// Get a specific event details.
func (app App) GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var event Event
	notFoundError := app.Conn.Find(&event, params["event_id"]).Where("created_at > ?", time.Now()).Error //ToDO:: Need to exempt from the done events.
	if notFoundError == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(event)
	} else {
		// Record not found for the provided Id.
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Event not found for provided id " + params["event_id"])
	}
}

func (j *eventDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = eventDate(t)
	return nil
}

func (j eventDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
