package leankit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gwenker/screan/server/configuration"
	"github.com/gwenker/screan/server/models"
)

// getBoard call leankit api to get board
func getBoard(boardID string) BoardResponse {
	var conf configuration.Configuration
	conf = configuration.GetConfiguration()
	fmt.Println(conf)
	// Build the request
	req, err := http.NewRequest("GET", conf.LeankitBoardURL+boardID, nil)
	req.SetBasicAuth(conf.LeankitUsername, conf.LeankitPassword)

	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var board BoardResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&board); err != nil {
		log.Println(err)
	}

	return board
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetUserStories to get user stories from leankit
func GetUserStories(stream models.Stream) []models.UserStory {
	var userStories []models.UserStory

	for _, board := range stream.Boards {
		if board.Type == "USERSTORIES" {
			boardLeankit := getBoard(board.ID)

			var childsLaneID []int

			for _, lane := range boardLeankit.ReplyData[0].Lanes {
				if board.LaneName != "" && lane.Title == board.LaneName {
					for _, laneChildID := range lane.ChildLaneIds {
						childsLaneID = append(childsLaneID, laneChildID)
					}
				}
				if contains(childsLaneID, lane.ID) || board.LaneName == "" {
					if len(lane.ChildLaneIds) > 0 {
						for _, laneChildID := range lane.ChildLaneIds {
							childsLaneID = append(childsLaneID, laneChildID)
						}
					} else {
						for _, card := range lane.Cards {
							var userStory models.UserStory
							userStory.ID = card.ExternalCardID
							userStory.Name = card.Title
							userStory.Complexity = card.Size
							userStory.LeftDaysToDevelop = float64(card.DrillThroughProgressSizeTotal-card.DrillThroughProgressSizeComplete) / 8.0
							userStory.TotalDaysToDevelop = float64(card.DrillThroughProgressSizeTotal) / 8.0
							userStories = append(userStories, userStory)
						}
					}
				}
			}
		}
	}

	userStories = GetTaskForUserStories(userStories, stream)

	return userStories
}

// GetTaskForUserStories to get taks of user stories from leankit
func GetTaskForUserStories(userStories []models.UserStory, stream models.Stream) []models.UserStory {

	for _, board := range stream.Boards {
		if board.Type == "TASKS" {
			boardLeankit := getBoard(board.ID)

			var childsLaneID []int

			for _, lane := range boardLeankit.ReplyData[0].Backlog {
				log.Println(lane.Title)
				if board.LaneName != "" && lane.Title == board.LaneName {
					for _, laneChildID := range lane.ChildLaneIds {
						childsLaneID = append(childsLaneID, laneChildID)
					}
				}
				if contains(childsLaneID, lane.ID) || board.LaneName == "" {
					if len(lane.ChildLaneIds) > 0 {
						for _, laneChildID := range lane.ChildLaneIds {
							childsLaneID = append(childsLaneID, laneChildID)
						}
					} else {
						for _, card := range lane.Cards {
							for idxUs, userStory := range userStories {
								if userStory.ID == card.ExternalCardID {
									var task models.Task
									task.ID = card.ID
									task.Name = card.Title
									task.Description = card.Description
									task.LeftDaysToDevelop = float64(card.Size) / 8.0

									f64, err := strconv.ParseFloat(card.Tags, 64)
									if err != nil {
										log.Println("Impossible to parse tag", card.Tags, err)
									} else {
										task.TotalDaysToDevelop = f64 / 8.0
									}

									userStories[idxUs].Tasks = append(userStories[idxUs].Tasks, task)
									break
								}
							}
						}
					}
				}
			}

			for _, lane := range boardLeankit.ReplyData[0].Lanes {
				log.Println(lane.Title)
				if board.LaneName != "" && lane.Title == board.LaneName {
					for _, laneChildID := range lane.ChildLaneIds {
						childsLaneID = append(childsLaneID, laneChildID)
					}
				}
				if contains(childsLaneID, lane.ID) || board.LaneName == "" {
					if len(lane.ChildLaneIds) > 0 {
						for _, laneChildID := range lane.ChildLaneIds {
							childsLaneID = append(childsLaneID, laneChildID)
						}
					} else {
						for _, card := range lane.Cards {
							for idxUs, userStory := range userStories {
								if userStory.ID == card.ExternalCardID {
									var task models.Task
									task.ID = card.ID
									task.Name = card.Title
									task.Description = card.Description
									task.LeftDaysToDevelop = float64(card.Size) / 8.0

									f64, err := strconv.ParseFloat(card.Tags, 64)
									if err != nil {
										log.Println("Impossible to parse tag", card.Tags, err)
									} else {
										task.TotalDaysToDevelop = f64 / 8.0
									}

									userStories[idxUs].Tasks = append(userStories[idxUs].Tasks, task)
									break
								}
							}
						}
					}
				}
			}
		}
	}

	return userStories
}
