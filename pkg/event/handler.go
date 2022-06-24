package event

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	repo EventRepository
}

func NewTask(eveRepo EventRepository) *Task {
	return &Task{
		repo: eveRepo,
	}
}

// CreateEvent
func (eve *Task) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Event
	requestBodyError := json.NewDecoder(r.Body).Decode(&event)
	if requestBodyError != nil {
		// Error in the request body.
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(requestBodyError.Error())
		if err != nil {
			return
		}
		return
	}

	event, error := eve.repo.CreateEvent(event)
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
func (eve *Task) GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var event Event
	event, notFoundError := eve.repo.GetEvent(params["event_id"])
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

// GetUpcomingEvents
// Considering event for complete day, which starts at 00:00 hours
func (eve *Task) GetUpcomingEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event []Event
	event, error := eve.repo.GetUpcomingEvents()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(error.Error())
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(event)
	if err != nil {
		return
	}
	return

}