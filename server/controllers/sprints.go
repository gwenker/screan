package controllers

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gwenker/screan/server/models"
	"github.com/gwenker/screan/server/modules/db"
	"github.com/gwenker/screan/server/modules/leankit"
	"github.com/labstack/echo"
)

// Sprints controller struct
type Sprints struct{}

// CreateSprint e.POST("/sprints", CreateSprint)
func (sc Sprints) CreateSprint(c echo.Context) error {
	// Get db from echo context
	session, ok := c.Get("db").(*db.DB)
	if !ok {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "db service is unavailable",
		})
	}

	s := models.Sprint{}
	// Bind
	if err := c.Bind(&s); err != nil {
		return err
	}
	// Add an Id
	s.ID = bson.NewObjectId()
	// Add creation date
	s.Created = time.Now()
	// Write the sprint to mongo
	_, err := session.Sprints().Create(s)
	if err != nil {
		return c.JSON(http.StatusFailedDependency, map[string]string{
			"message": err.Error(),
		})
	}
	// Return sprint created
	return c.JSON(http.StatusCreated, s)
}

// GetSprints e.GET("/sprints", GetSprints)
func (sc Sprints) GetSprints(c echo.Context) error {
	// Get db from echo context
	session, ok := c.Get("db").(*db.DB)
	if !ok {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "db service is unavailable",
		})
	}

	sprints, err := session.Sprints().FindAll()
	if err != nil {
		return c.JSON(http.StatusFailedDependency, map[string]string{
			"message": err.Error(),
		})
	}

	// Return Sprints found
	return c.JSON(http.StatusOK, sprints)
}

// GetSprint e.GET("/sprints/:id", GetSprint)
func (sc Sprints) GetSprint(c echo.Context) error {
	id := c.Param("id")
	// Get db from echo context
	session, ok := c.Get("db").(*db.DB)
	if !ok {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "db service is unavailable",
		})
	}

	sprint, err := session.Sprints().FindByID(id)
	if err != nil {
		return err
	}

	sprint.UserStories = leankit.GetUserStories(sprint.Stream)

	// Return Sprint found
	return c.JSON(http.StatusOK, sprint)
}

// UpdateSprint e.update("/sprints/:id", UpdateSprint)
func (sc Sprints) UpdateSprint(c echo.Context) error {
	id := c.Param("id")
	// Get db from echo context
	session, ok := c.Get("db").(*db.DB)
	if !ok {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "db service is unavailable",
		})
	}

	// Updated sprint
	s := models.Sprint{}
	// Bind
	if err := c.Bind(&s); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if _, err := session.Sprints().Update(id, s); err != nil {
		return c.JSON(http.StatusFailedDependency, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, s)
}

// RemoveSprint e.DELETE("/sprints/:id", RemoveSprint)
func (sc Sprints) RemoveSprint(c echo.Context) error {
	id := c.Param("id")
	// Get db from echo context
	session, ok := c.Get("db").(*db.DB)
	if !ok {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "db service is unavailable",
		})
	}

	// Remove sprint
	if err := session.Sprints().Delete(id); err != nil {
		return c.JSON(http.StatusFailedDependency, map[string]string{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
