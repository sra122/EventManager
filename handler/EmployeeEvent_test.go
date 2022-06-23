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
	reader := strings.NewReader("{\n    \"employee_id\" : 1\n}")
	req := httptest.NewRequest(http.MethodPost, "/event/{event_id}/employees", reader)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"event_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	h.AddEmployeeForEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandler_AddEmployeeForEventWithoutRequestBody(t *testing.T) {
	h := initialise()
	req := httptest.NewRequest(http.MethodPost, "/event/{event_id}/employees", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"event_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	h.AddEmployeeForEvent(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_GetEmployeesForEvent(t *testing.T) {
	h := initialise()
	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}/employees", nil)
	w := httptest.NewRecorder()
	h.GetEmployeesForEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandler_GetEmployeesForEventWithAccommodationQuery(t *testing.T) {
	h := initialise()
	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}/employees", nil)
	req.URL.Query().Set("is_accommodation_required", "true")
	w := httptest.NewRecorder()
	h.GetEmployeesForEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
