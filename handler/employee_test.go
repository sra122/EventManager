package handler

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateEmployeeWithWrongRequestBody(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : 123,\n    \"email\" : \"sravan@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	dropTable(h)
}

func TestHandler_CreateEmployeeWithWrongDateFormat(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"02-01-2006\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"sravan@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	dropTable(h)
}

func TestHandler_CreateEmployeeWithSameEmailId(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"02-01-2006\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"sravan@gmail.com\" \n}")
	httptest.NewRequest(http.MethodPost, "/employees", reader)
	request := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, request)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	dropTable(h)
}

func TestHandler_CreateEmployeeSuccessScenario(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	dropTable(h)
}

func TestHandler_UpdateEmployeeSuccessScenario(t *testing.T) {
	h := initialise()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-03\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"testing@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPut, "/employees/{employee_id}", reader)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	dropTable(h)
}

func TestHandler_GetEmployees(t *testing.T) {
	h := initialise()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	w := httptest.NewRecorder()
	h.GetEmployees(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	dropTable(h)
}
