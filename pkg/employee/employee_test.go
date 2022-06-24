package employee

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateEmployeeWithWrongRequestBody(t *testing.T) {
	h, conn := InitializeDBConnection()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : 123,\n    \"email\" : \"sravan@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	DropTable(*conn)
}

func TestHandler_CreateEmployeeWithWrongDateFormat(t *testing.T) {
	h, conn := InitializeDBConnection()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"02-01-2006\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"sravan@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	DropTable(*conn)
}

func TestHandler_CreateEmployeeWithSameEmailId(t *testing.T) {
	h, conn := InitializeDBConnection()
	readerFirst := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"sravan@gmail.com\" \n}")
	firstRequest := httptest.NewRequest(http.MethodPost, "/employees", readerFirst)
	firstWrite := httptest.NewRecorder()
	h.CreateEmployee(firstWrite, firstRequest)
	assert.Equal(t, http.StatusCreated, firstWrite.Code)

	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"sravan@gmail.com\" \n}")
	request := httptest.NewRequest(http.MethodPost, "/employees", reader)
	write := httptest.NewRecorder()
	h.CreateEmployee(write, request)
	assert.Equal(t, http.StatusInternalServerError, write.Code)
	DropTable(*conn)
}

func TestHandler_CreateEmployeeSuccessScenario(t *testing.T) {
	h, conn := InitializeDBConnection()
	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	w := httptest.NewRecorder()
	h.CreateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	DropTable(*conn)
}

func TestHandler_UpdateEmployeeWithInvalidEmployeeId(t *testing.T) {
	h, conn := InitializeDBConnection()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-03\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"testing@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPut, "/employees/{employee_id}", reader)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "10",
	}
	req = mux.SetURLVars(req, vars)

	h.UpdateEmployee(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	DropTable(*conn)
}

func TestHandler_UpdateEmployeeWithBadBodyRequest(t *testing.T) {
	h, conn := InitializeDBConnection()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	reader := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"02-01-2000\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"testing@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPut, "/employees/{employee_id}", reader)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	h.UpdateEmployee(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	DropTable(*conn)
}

func TestHandler_UpdateEmployeeSuccessScenario(t *testing.T) {
	h, conn := InitializeDBConnection()

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

	h.UpdateEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	DropTable(*conn)
}

func TestHandler_UpdateEmployeeWithNullAsFirstName(t *testing.T) {
	h, conn := InitializeDBConnection()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	reader := strings.NewReader("{\n    \"firstName\": null,\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-03\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"testing@gmail.com\" \n}")
	req := httptest.NewRequest(http.MethodPut, "/employees/{employee_id}", reader)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	h.UpdateEmployee(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	DropTable(*conn)
}

func TestHandler_GetEmployees(t *testing.T) {
	h, conn := InitializeDBConnection()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	w := httptest.NewRecorder()
	h.GetEmployees(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	DropTable(*conn)
}

func TestHandler_DeleteEmployeeWithInvalidEmployeeId(t *testing.T) {
	h, conn := InitializeDBConnection()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	req := httptest.NewRequest(http.MethodDelete, "/employees/{employee_id}", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "10",
	}
	req = mux.SetURLVars(req, vars)

	h.DeleteEmployee(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	DropTable(*conn)
}

func TestHandler_DeleteEmployeeSuccessScenario(t *testing.T) {
	h, conn := InitializeDBConnection()

	readerEmployeeCreate := strings.NewReader("{\n    \"firstName\": \"Sravan\",\n    \"lastName\" : \"Hello\",\n    \"birthDay\" : \"2006-01-02\",\n    \"gender\" : \"male\",\n    \"is_accommodation_required\" : true,\n    \"email\" : \"test@gmail.com\" \n}")
	reqEmployee := httptest.NewRequest(http.MethodPost, "/employees", readerEmployeeCreate)
	writeEmployee := httptest.NewRecorder()
	h.CreateEmployee(writeEmployee, reqEmployee)
	assert.Equal(t, http.StatusCreated, writeEmployee.Code)

	req := httptest.NewRequest(http.MethodDelete, "/employees/{employee_id}", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"employee_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	h.DeleteEmployee(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	DropTable(*conn)
}
