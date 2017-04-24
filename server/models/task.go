package models

// Task is a part of a feature to develop in a sprint
type Task struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	TotalDaysToDevelop float64 `json:"totalDaysToDevelop"`
	LeftDaysToDevelop  float64 `json:"leftDaysToDevelop"`
}
