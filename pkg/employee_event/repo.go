package employee_event

import (
	"example.com/hello/pkg/event"
	"gorm.io/gorm"
)

// EventEmployees Model
type EventEmployees struct {
	EmployeeId int `json:"employee_id" gorm:"not null"`
	EventId    int `json:"event_id"`
}

// Handler
type EmployeeEventConnection struct {
	DB *gorm.DB
}

func InitialiseEmployeeEventHandler(db *gorm.DB) *EmployeeEventConnection {
	return &EmployeeEventConnection{DB: db}
}

type EmployeeEventRepository interface {
	AddEmployeeToEvent(eventEmployees EventEmployees) (EventEmployees, error)
	GetEmployeesForEvent(eventId string) (event.Event, error)
	FindEvent(eventId string) (event.Event, error)
	GetEmployeesForEventWithAccommodationQuery(eventId string, accommodationQuery string) (event.Event, error)
}

func (h *EmployeeEventConnection) AddEmployeeToEvent(eventEmployees EventEmployees) (EventEmployees, error) {
	err := h.DB.Create(&eventEmployees).Error

	if err != nil {
		return EventEmployees{}, err
	}
	return eventEmployees, nil
}

func (h *EmployeeEventConnection) FindEvent(eventId string) (event.Event, error) {
	var eve event.Event
	notFoundError := h.DB.First(&eve, eventId).Error
	if notFoundError != nil {
		return event.Event{}, notFoundError
	}
	return eve, nil
}

func (h *EmployeeEventConnection) GetEmployeesForEvent(eventId string) (event.Event, error) {
	var eve event.Event
	err := h.DB.Preload("Employees").Find(&eve, eventId).Error
	if err != nil {
		return event.Event{}, err
	}
	return eve, nil
}

func (h *EmployeeEventConnection) GetEmployeesForEventWithAccommodationQuery(eventId string, accommodationQuery string) (event.Event, error) {
	var eve event.Event
	err := h.DB.Preload("Employees", "is_accommodation_required", accommodationQuery).Find(&eve, eventId).Error
	if err != nil {
		return event.Event{}, err
	}
	return eve, nil
}
