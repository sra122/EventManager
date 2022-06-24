package employee

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Task struct {
	repo EmployeeRepository
}

func NewTask(empRepo EmployeeRepository) *Task {
	return &Task{
		repo: empRepo,
	}
}

// GetEmployees
// Get List of Employees in the Organization/**
func (emp *Task) GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee, error := emp.repo.GetEmployees()
	if error == nil {
		err := json.NewEncoder(w).Encode(employee)
		if err != nil {
			return
		}
	}
}

// CreateEmployee
// Create an Employee to the Organization
func (emp *Task) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	requestBodyError := json.NewDecoder(r.Body).Decode(&employee)
	if requestBodyError != nil {
		// If request body doesn't fit according to the requirements
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(strings.Trim(requestBodyError.Error(), "\""))
		if err != nil {
			return
		}
		return
	}
	employee, error := emp.repo.CreateEmployee(employee)
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(employee)
		if err != nil {
			return
		}
	} else {
		//Error occurred during creating of employee
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(error.Error())
		if err != nil {
			return
		}
	}
}

// UpdateEmployee
// Updates the Employee for the provided id in the url
func (emp *Task) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	employee, notFoundError := emp.repo.FetchFirstEmployee(employee, params["employee_id"])
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Employee record is not found with id " + params["employee_id"])
		if err != nil {
			return
		}
		return
	}
	requestBodyError := json.NewDecoder(r.Body).Decode(&employee)
	if requestBodyError != nil {
		// If request body doesn't fit according to the requirements
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(strings.Trim(requestBodyError.Error(), "\""))
		if err != nil {
			return
		}
		return
	}
	employee, error := emp.repo.UpdateEmployee(employee)
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(employee)
		if err != nil {
			return
		}
	} else {
		// Error occurred during the save of Entity
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(error.Error())
		if err != nil {
			return
		}
	}
}

// DeleteEmployee
// Delete Employee from the records
func (emp *Task) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	employee, notFoundError := emp.repo.FetchFirstEmployee(employee, params["employee_id"])
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Employee record is not found with id " + params["employee_id"])
		if err != nil {
			return
		}
		return
	}

	employee, error := emp.repo.DeleteEmployee(employee, params["employee_id"])
	if error == nil {
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode("Employee is deleted ||")
		if err != nil {
			return
		}
	} else {
		// Error occurred during the delete operation.
		w.WriteHeader(http.StatusBadGateway)
		err := json.NewEncoder(w).Encode(error.Error())
		if err != nil {
			return
		}
	}
}
