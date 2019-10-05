package model

import (
	"time"
)

type Adventure struct {
	Type string	`firestore:"type"`
	Name string	`firestore:"name"`
	Location string	`firestore:"location"`
	EstimatedDuration string `firestore:"estimatedDuration"`
	Price	float64	`firestore:"price"`
	CreatedAt	time.Time	`firestore:"createdAt"`
	Status	Status	`firestore:"status"`
}

func NewAdventure() *Adventure {
	return &Adventure{
		CreatedAt: time.Now(),
	}
}

func AdventureToResponseAdventureCreate(a *Adventure) *ResponseAdventureCreate {
	if a == nil {
		return &ResponseAdventureCreate{}
	}

	return &ResponseAdventureCreate{
		Type: a.Type,
		Name: a.Name,
		Location: a.Location,
		EstimatedDuration: a.EstimatedDuration,
		Price: a.Price,
		CreatedAt: a.CreatedAt,
		Status: a.Status,
	}
}