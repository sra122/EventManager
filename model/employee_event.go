package model

type EventEmployees struct {
	EmployeeId int `json:"employee_id" gorm:"not null"`
	EventId    int `json:"event_id"`
}
