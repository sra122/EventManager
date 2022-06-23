package model

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
