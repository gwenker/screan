package server

import (
	"strconv"

	"github.com/gwenker/screan/server/sprint"
	"github.com/labstack/echo"
)

// Start Start server
func Start(port int) {
	e := echo.New()
	e.Debug = true
	e.POST("/sprints", sprint.CreateSprint)
	e.GET("/sprints", sprint.GetSprints)
	e.GET("/sprints/:id", sprint.GetSprint)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}
