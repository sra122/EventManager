package handler

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_AddEmployeeForEvent(t *testing.T) {
	h := initialise()

	readerEmployee := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployee)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	readerEvent := strings.NewReader("{\n    \"name\" : \"New Year\",\n    \"date\" : \"2022-12-31\"\n}")
	reqEvent := httptest.NewRequest(http.MethodPost, "/event", readerEvent)
	writeEvent := httptest.NewRecorder()
	h.CreateEvent(writeEvent, reqEvent)
	assert.Equal(t, http.StatusCreated, writeEvent.Code)

	readerEmployeeEvent := strings.NewReader("{\n    \"employee_id\" : 1\n}")
	reqEmployeeEvent := httptest.NewRequest(http.MethodPost, "/event/{event_id}/employees", readerEmployeeEvent)
	writeEmployeeEvent := httptest.NewRecorder()
	vars := map[string]string{
		"event_id": "1",
	}
	reqEmployeeEvent = mux.SetURLVars(reqEmployeeEvent, vars)
	h.AddEmployeeForEvent(writeEmployeeEvent, reqEmployeeEvent)
	assert.Equal(t, http.StatusCreated, writeEmployeeEvent.Code)

	dropTable(h)
}

func TestHandler_AddEmployeeForEventWithoutRequestBody(t *testing.T) {
	h := initialise()
	readerEmployee := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployee)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	readerEvent := strings.NewReader("{\n    \"name\" : \"New Year\",\n    \"date\" : \"2022-12-31\"\n}")
	reqEvent := httptest.NewRequest(http.MethodPost, "/event", readerEvent)
	writeEvent := httptest.NewRecorder()
	h.CreateEvent(writeEvent, reqEvent)
	assert.Equal(t, http.StatusCreated, writeEvent.Code)

	req := httptest.NewRequest(http.MethodPost, "/event/{event_id}/employees", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"event_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	h.AddEmployeeForEvent(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	dropTable(h)
}

func TestHandler_GetEmployeesForEvent(t *testing.T) {
	h := initialise()

	readerEmployee := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployee)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	readerEvent := strings.NewReader("{\n    \"name\" : \"New Year\",\n    \"date\" : \"2022-12-31\"\n}")
	reqEvent := httptest.NewRequest(http.MethodPost, "/event", readerEvent)
	writeEvent := httptest.NewRecorder()
	h.CreateEvent(writeEvent, reqEvent)
	assert.Equal(t, http.StatusCreated, writeEvent.Code)

	readerEmployeeEvent := strings.NewReader("{\n    \"employee_id\" : 1\n}")
	reqEmployeeEvent := httptest.NewRequest(http.MethodPost, "/event/{event_id}/employees", readerEmployeeEvent)
	writeEmployeeEvent := httptest.NewRecorder()
	vars := map[string]string{
		"event_id": "1",
	}
	reqEmployeeEvent = mux.SetURLVars(reqEmployeeEvent, vars)
	h.AddEmployeeForEvent(writeEmployeeEvent, reqEmployeeEvent)
	assert.Equal(t, http.StatusCreated, writeEmployeeEvent.Code)

	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}/employees", nil)
	w := httptest.NewRecorder()
	h.GetEmployeesForEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	dropTable(h)
}

func TestHandler_GetEmployeesForEventWithAccommodationQuery(t *testing.T) {
	h := initialise()

	readerEmployee := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployee)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	readerEvent := strings.NewReader("{\n    \"name\" : \"New Year\",\n    \"date\" : \"2022-12-31\"\n}")
	reqEvent := httptest.NewRequest(http.MethodPost, "/event", readerEvent)
	writeEvent := httptest.NewRecorder()
	h.CreateEvent(writeEvent, reqEvent)
	assert.Equal(t, http.StatusCreated, writeEvent.Code)

	readerEmployeeEvent := strings.NewReader("{\n    \"employee_id\" : 1\n}")
	reqEmployeeEvent := httptest.NewRequest(http.MethodPost, "/event/{event_id}/employees", readerEmployeeEvent)
	writeEmployeeEvent := httptest.NewRecorder()
	vars := map[string]string{
		"event_id": "1",
	}
	reqEmployeeEvent = mux.SetURLVars(reqEmployeeEvent, vars)
	h.AddEmployeeForEvent(writeEmployeeEvent, reqEmployeeEvent)
	assert.Equal(t, http.StatusCreated, writeEmployeeEvent.Code)

	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}/employees", nil)
	req.URL.Query().Set("is_accommodation_required", "true")
	w := httptest.NewRecorder()
	h.GetEmployeesForEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	dropTable(h)

}
