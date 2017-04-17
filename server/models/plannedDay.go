package models

import "time"

// PlannedDay define a planned day
type PlannedDay struct {
	Day                time.Time `json:"day" bson:"day"`
	DayCapacity        float32   `json:"dayCapacity" bson:"dayCapacity"`
	TheoricalPointToDo float32   `json:"theoricalPointToDo" bson:"theoricalPointToDo"`
}
