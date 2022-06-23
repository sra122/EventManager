package handler

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_GetEventSuccessScenario(t *testing.T) {
	h := initialise()
	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"event_id": "1",
	}
	req = mux.SetURLVars(req, vars)
	h.GetEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	dropTable(h)
}

func TestHandler_GetEventWithInvalidEventId(t *testing.T) {
	h := initialise()
	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"event_id": "abc",
	}
	req = mux.SetURLVars(req, vars)
	h.GetEvent(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	dropTable(h)
}

func TestHandler_CreateEvent(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"name\" : \"New Year\",\n    \"date\" : \"2022-12-31\"\n}")
	req := httptest.NewRequest(http.MethodPost, "/event", reader)
	w := httptest.NewRecorder()

	h.CreateEvent(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	dropTable(h)
}

func TestHandler_CreateEventForWrongDateFormat(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"name\" : \"New Year\",\n    \"date\" : \"12-12-2022\"\n}")
	req := httptest.NewRequest(http.MethodPost, "/event", reader)
	w := httptest.NewRecorder()

	h.CreateEvent(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	dropTable(h)
}

func TestHandler_CreateEventWithWrongRequestBody(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n\"date\" : \"12-12-2022\"\n}")
	req := httptest.NewRequest(http.MethodPost, "/event", reader)
	w := httptest.NewRecorder()

	h.CreateEvent(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	dropTable(h)
}

func TestHandler_CreateEventWithNullTitle(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n    \"name\" : null,\n    \"date\" : \"2022-07-07\"\n}")
	req := httptest.NewRequest(http.MethodPost, "/event", reader)
	w := httptest.NewRecorder()

	h.CreateEvent(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	dropTable(h)
}

func TestHandler_GetUpcomingEvents(t *testing.T) {
	h := initialise()
	reader := strings.NewReader("{\n\"date\" : \"12-12-2022\"\n}")
	req := httptest.NewRequest(http.MethodPost, "/event", reader)
	w := httptest.NewRecorder()

	h.GetUpcomingEvents(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	dropTable(h)
}
