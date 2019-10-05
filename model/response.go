package model

import(
	"time"
)

type Status string

const (
	StatusCompleted	Status = "completed"
	StatusRunning Status = "running"
	StatusError Status = "error"
)

type ResponseAdventureCreate struct {
	Type string	`json:"type,omitempty"`
	Name string	`json:"name,omitempty"`
	Location string	`json:"location,omitempty"`
	EstimatedDuration string `json:"estimatedDuration,omitempty"`
	Price	float64	`json:"price,omitempty"`
	CreatedAt	time.Time	`json:"createdAt"`
	Status	Status	`json:"status,omitempty"`
}
