package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Configuration struct
type Configuration struct {
	LeankitUsername string
	LeankitPassword string
	LeankitBoardURL string
}

// GetConfiguration to load conf file and access to configuration properties
func GetConfiguration() Configuration {
	file, errFile := os.Open("./server/configuration/conf.json")
	if errFile != nil {
		log.Fatal(errFile)
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("config:", configuration)
	return configuration
}
