package server

import (
	"strconv"

	mgo "gopkg.in/mgo.v2"

	"github.com/gwenker/screan/server/sprint"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start Start server
func Start(port int) {
	e := echo.New()
	e.Debug = true

	// FIXME
	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())

	// Get mongodb session
	ms := getSession()
	defer ms.Close()

	// Get a SprintController instance
	sc := sprint.New(ms)

	// routes
	e.POST("/sprints", sc.CreateSprint)
	e.GET("/sprints", sc.GetSprints)
	e.GET("/sprints/:id", sc.GetSprint)
	e.PUT("/sprints/:id", sc.UpdateSprint)
	e.DELETE("/sprints/:id", sc.RemoveSprint)

	// start echo server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}

// get mongodb session
func getSession() *mgo.Session {
	// Connect to local mongo
	session, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return session
}
