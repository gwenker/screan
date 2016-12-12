package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Sprint define an iteration of development
type Sprint struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Created   time.Time     `json:"created" bson:"created"`
	Name      string        `json:"name" bson:"name"`
	StartDate time.Time     `json:"startDate" bson:"startDate"`
	EndDate   time.Time     `json:"endDate" bson:"endDate"`
	Stream    Stream        `json:"stream" bson:"stream"`
	Capacity  float32       `json:"capacity" bson:"capacity"`
	Velocity  float32       `json:"velocity" bson:"velocity"`
	Planning  []PlannedDay  `json:"planning" bson:"planning"`
	Events    []Event       `json:"events" bson:"events"`
}
