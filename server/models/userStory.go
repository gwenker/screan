package models

// UserStory is feature to develop in a sprint
type UserStory struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Complexity         string `json:"complexity"`
	TotalDaysToDevelop string `json:"totalDaysToDevelop"`
	LeftDaysToDevelop  string `json:"leftDaysToDevelop"`
	State              string `json:"state"`
}
