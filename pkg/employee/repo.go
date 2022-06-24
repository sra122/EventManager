package employee

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
	"time"
)

type birthDay time.Time

type Employee struct {
	gorm.Model
	FirstName               *string  `json:"firstName" gorm:"not null"`
	LastName                *string  `json:"lastName" gorm:"not null"`
	BirthDay                birthDay `json:"birthDay"`
	Gender                  string   `json:"gender"` // Gender
	IsAccommodationRequired bool     `json:"is_accommodation_required"`
	Email                   *string  `json:"email" gorm:"unique;not null"`
}

// Handler
type EmployeeConnection struct {
	DB *gorm.DB
}

func InitialiseEmployeeHandler(db *gorm.DB) *EmployeeConnection {
	return &EmployeeConnection{DB: db}
}

// Interface
type EmployeeRepository interface {
	GetEmployees() ([]Employee, error)
	CreateEmployee(employee Employee) (Employee, error)
	UpdateEmployee(employee Employee) (Employee, error)
	FetchFirstEmployee(employee Employee, employeeId string) (Employee, error)
	DeleteEmployee(employee Employee, employeeId string) (Employee, error)
}

func (h *EmployeeConnection) GetEmployees() ([]Employee, error) {
	var employee []Employee
	e := h.DB.Order("id asc").Find(&employee).Error
	if e != nil {
		return nil, e
	}
	return employee, nil
}

func (h *EmployeeConnection) CreateEmployee(employee Employee) (Employee, error) {
	error := h.DB.Create(&employee).Error

	if error != nil {
		return Employee{}, error
	}
	return employee, nil
}

func (h *EmployeeConnection) FetchFirstEmployee(employee Employee, employeeId string) (Employee, error) {
	notFoundError := h.DB.First(&employee, employeeId).Error

	if notFoundError != nil {
		return Employee{}, notFoundError
	}
	return employee, nil
}

func (h *EmployeeConnection) UpdateEmployee(employee Employee) (Employee, error) {
	error := h.DB.Save(&employee).Error

	if error != nil {
		return Employee{}, error
	}

	return employee, nil
}

func (h *EmployeeConnection) DeleteEmployee(employee Employee, employeeId string) (Employee, error) {
	error := h.DB.Delete(&employee, employeeId).Error

	if error != nil {
		return Employee{}, error
	}
	return employee, nil
}

func (j *birthDay) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = birthDay(t)
	return nil
}

func (j birthDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
