package models

// Board is a board with cards UserStories or DevelopmentTasks
type Board struct {
	ID       string `json:"id"`
	LaneName string `json:"laneName"`
	Type     string `json:"type"`
}
