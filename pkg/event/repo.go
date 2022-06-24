package event

import (
	"encoding/json"
	"example.com/hello/dbconnection"
	"example.com/hello/model"
	"example.com/hello/pkg/employee"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
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

func InitializeDBConnection() (*Task, *EventConnection) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB := dbconnection.ConnectTestDb()

	empRepo := InitialiseEventHandler(DB)
	empTask := NewTask(empRepo)

	return empTask, empRepo
}

// Drop the table after testing.
func DropTable(conn EventConnection) {
	conn.DB.Migrator().DropTable(&model.Event{})
}

func (h *EventConnection) CreateEvent(event Event) (Event, error) {
	error := h.DB.Create(&event).Error

	if error != nil {
		return Event{}, nil
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
