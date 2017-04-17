package leankit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const boardURL = "https://*********.leankit.com/kanban/api/boards/*********"

// GetBoard call leankit api to get board
func GetBoard() {
	// Build the request
	req, err := http.NewRequest("GET", boardURL, nil)
	req.SetBasicAuth("*********", "************")
    req.
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
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
		return
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

	fmt.Println(board)
}
