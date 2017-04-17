package models

// Stream define a team board
type Stream struct {
	StreamName string `json:"name" bson:"name"`
	BoardID    string `json:"boardId" bson:"created"`
	LaneName   string `json:"laneName" bson:"laneName"`
}
