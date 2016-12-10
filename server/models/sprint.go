package models

import "gopkg.in/mgo.v2/bson"

type Sprint struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	StartDate string        `json:"startDate" bson:"startDate"`
	EndDate   string        `json:"endDate" bson:"endDate"`
}
