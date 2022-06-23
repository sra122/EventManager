package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
	"time"
)

type birthDay time.Time

type Employee struct {
	gorm.Model                       //TODO: Add email for converting unique record.
	FirstName               string   `json:"firstName"`
	LastName                string   `json:"lastName"`
	BirthDay                birthDay `json:"birthDay"`
	Gender                  string   `json:"gender"` // Gender
	IsAccommodationRequired bool     `json:"is_accommodation_required"`
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
