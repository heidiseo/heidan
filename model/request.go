package model

type RequestAdventureCreate struct {
	Type string	`json:"type,omitempty"`
	Name string	`json:"name,omitempty"`
	Location string	`json:"location,omitempty"`
	EstimatedDuration string `json:"estimatedDuration,omitempty"`
	Price	float64	`json:"price,omitempty"`
}

func RequestAdventureCreatetoAdventure(reqAdventure *RequestAdventureCreate) *Adventure {
	if reqAdventure == nil {
		return NewAdventure()
	}
	a := NewAdventure()
	a.Type = reqAdventure.Type
	a.Name = reqAdventure.Name
	a.Location = reqAdventure.Location
	a.EstimatedDuration = reqAdventure.EstimatedDuration
	a.Price	= reqAdventure.Price
	return a
}
