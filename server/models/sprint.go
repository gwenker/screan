package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Sprint struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Created   time.Time     `json:"created:" bson:"created"`
	Name      string        `json:"name" bson:"name"`
	StartDate time.Time     `json:"startDate" bson:"startDate"`
	EndDate   time.Time     `json:"endDate" bson:"endDate"`
}
