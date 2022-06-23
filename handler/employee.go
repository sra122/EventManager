package handler

import (
	"encoding/json"
	"example.com/hello/model"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// GetEmployees
// Get List of Employees in the Organization/**
func (h Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee []model.Employee
	h.DB.Order("id asc").Find(&employee)
	err := json.NewEncoder(w).Encode(employee)
	if err != nil {
		return
	}
}

// CreateEmployee
// Create an Employee to the Organization
func (h Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee model.Employee
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
	error := h.DB.Create(&employee).Error
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
func (h Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee model.Employee
	notFoundError := h.DB.First(&employee, params["employee_id"]).Error
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
	error := h.DB.Save(&employee).Error
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
func (h Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee model.Employee
	notFoundError := h.DB.First(&employee, params["employee_id"]).Error
	if notFoundError != nil {
		//If employee record not found with the provided id.
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Employee record is not found with id " + params["employee_id"])
		if err != nil {
			return
		}
		return
	}

	error := h.DB.Delete(&employee, params["employee_id"]).Error
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
