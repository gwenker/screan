package models

import "time"

// PlannedDay define a planned day
type PlannedDay struct {
	Day                time.Time `json:"day" bson:"day"`
	DayCapacity        string    `json:"dayCapacity" bson:"dayCapacity"`
	TheoricalPointToDo int16     `json:"theoricalPointToDo" bson:"theoricalPointToDo"`
}
