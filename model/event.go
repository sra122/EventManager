package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
	"time"
)

type eventDate time.Time

type Event struct {
	gorm.Model
	Name      *string    `json:"name" gorm:"unique;not null"`
	Date      *eventDate `json:"date"`
	Employees []Employee `gorm:"many2many:event_employees;"`
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
