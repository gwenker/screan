package models

import "time"

// Event define an event
type Event struct {
	EventDate time.Time `json:"eventDate" bson:"eventDate"`
	Name      string    `json:"name" bson:"name"`
	Duration  float32   `json:"duration" bson:"duration"`
}
