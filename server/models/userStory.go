package models

// UserStory is feature to develop in a sprint
type UserStory struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Complexity         int     `json:"complexity"`
	TotalDaysToDevelop float64 `json:"totalDaysToDevelop"`
	LeftDaysToDevelop  float64 `json:"leftDaysToDevelop"`
	State              string  `json:"state"`
}
