package event

import (
	"encoding/json"
	"example.com/hello/pkg/employee"
	"gorm.io/gorm"
	"strings"
	"time"
)

type eventDate time.Time

// Event Model
type Event struct {
	gorm.Model
	Name      *string             `json:"name" gorm:"unique;not null"`
	Date      *eventDate          `json:"date"`
	Employees []employee.Employee `gorm:"many2many:event_employees;"`
}

// Handler
type EventConnection struct {
	DB *gorm.DB
}

func InitialiseEventHandler(db *gorm.DB) *EventConnection {
	return &EventConnection{DB: db}
}

type EventRepository interface {
	CreateEvent(event Event) (Event, error)
	GetEvent(eventId string) (Event, error)
	GetUpcomingEvents() ([]Event, error)
}

func (h *EventConnection) CreateEvent(event Event) (Event, error) {
	error := h.DB.Create(&event).Error

	if error != nil {
		return Event{}, error
	}
	return event, nil
}

func (h *EventConnection) GetEvent(eventId string) (Event, error) {
	var event Event
	notFoundError := h.DB.Find(&event, eventId).Error

	if notFoundError != nil {
		return Event{}, notFoundError
	}
	return event, nil
}

func (h *EventConnection) GetUpcomingEvents() ([]Event, error) {
	var events []Event
	err := h.DB.Order("id asc").Find(&events, "date > ?", time.Now().Add(-24*time.Hour)).Error

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (j *eventDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = eventDate(t)
	return nil
}

func (j eventDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
