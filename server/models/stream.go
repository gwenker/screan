package models

// Stream define a team board
type Stream struct {
	StreamName string  `json:"name" bson:"name"`
	Boards     []Board `json:"boards" bson:"boards"`
}
