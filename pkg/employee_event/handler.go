package employee_event

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Task struct {
	repo EmployeeEventRepository
}

func NewTask(empEvenRepo EmployeeEventRepository) *Task {
	return &Task{
		repo: empEvenRepo,
	}
}

// AddEmployeeForEvent
// Add an employee to the event.
func (empEve *Task) AddEmployeeForEvent(w http.ResponseWriter, r *http.Request) {
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
	eventEmployees, err := empEve.repo.AddEmployeeToEvent(eventEmployees)
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
func (empEve *Task) GetEmployeesForEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	accommodationQuery := r.URL.Query().Get("is_accommodation_required")
	var err error

	evn, notFoundError := empEve.repo.FindEvent(params["event_id"])
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Event record is not found with id " + params["event_id"])
		return
	}

	// If Query Parameters is not empty
	if accommodationQuery != "" {
		evn, err = empEve.repo.GetEmployeesForEventWithAccommodationQuery(params["event_id"], accommodationQuery)
	} else {
		evn, err = empEve.repo.GetEmployeesForEvent(params["event_id"])
	}

	if err == nil {
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(evn)
		if err != nil {
			return
		}
		return
	} else {
		// Error while fetching the records
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(err.Error())
		if err != nil {
			return
		}
		return
	}
}
