package server

import (
	"strconv"

	"github.com/gwenker/screan/server/controllers"
	"github.com/gwenker/screan/server/modules/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start Start server
func Start(port int) {
	engine := echo.New()
	engine.Debug = true

	// FIXME
	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	engine.Use(middleware.CORS())

	// Get mongodb session
	db.Connect("mongodb://localhost")
	defer db.Close()

	// Get a SprintController instance
	sc := controllers.Sprints{}

	// routes
	sprints := engine.Group("/sprints")
	{
		sprints.Use(db.Middleware("", ""))
		sprints.POST("", sc.CreateSprint)
		sprints.GET("", sc.GetSprints)
		sprints.GET("/:id", sc.GetSprint)
		sprints.PUT("/:id", sc.UpdateSprint)
		sprints.DELETE("/:id", sc.RemoveSprint)
	}
	// start echo server
	engine.Logger.Fatal(engine.Start(":" + strconv.Itoa(port)))
}
