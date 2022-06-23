package handler

import (
	"encoding/json"
	"example.com/hello/model"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// CreateEvent
func (h handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event model.Event
	requestBodyError := json.NewDecoder(r.Body).Decode(&event)
	if requestBodyError != nil {
		// Error in the requestbody.
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(requestBodyError.Error())
		if err != nil {
			return
		}
		return
	}

	error := h.DB.Create(&event).Error
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(event)
		if err != nil {
			return
		}
	} else {
		// Error while creating an entity
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(error.Error())
		if err != nil {
			return
		}
	}
}

// GetEvent
// Get a specific event details.
func (h handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var event model.Event
	notFoundError := h.DB.Find(&event, params["event_id"]).Where("created_at > ?", time.Now().Add(-24*time.Hour)).Error //TODO:: Need to exempt from the done events.
	if notFoundError == nil {
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(event)
		if err != nil {
			return
		}
	} else {
		// Record not found for the provided Id.
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Event not found for provided id " + params["event_id"])
		if err != nil {
			return
		}
	}
}
