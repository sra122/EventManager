package handler

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_GetEmployees(t *testing.T) {
	h := initialise()
	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	w := httptest.NewRecorder()
	h.GetEmployees(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandler_CreateEmployeeSuccessScenario(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\" : \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : false\n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// "is_accommodation_required" variable is passed as string/integer instead of Boolean.
func TestHandler_CreateEmployeeWithWrongRequestBody(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\" : \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : 123\n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_CreateEmployeeWithWrongDateFormat(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\" : \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"21-01-2006\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : 123\n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_UpdateEmployeeSuccessScenario(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\" : \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : false\n}")
	req := httptest.NewRequest(http.MethodPut, "/employees/{employee_id}", reader)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
