package handler

import (
	"example.com/hello/dbconnection"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func initialise() handler {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectDb()
	h := New(DB)

	return h
}

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
	req := httptest.NewRequest(http.MethodPut, "/employees/1", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
