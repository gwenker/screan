package sprint

import (
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gwenker/screan/server/models"
	"github.com/labstack/echo"
)

// Controller Sprint struct
type Controller struct {
	session *mgo.Session
}

// New create a new sprint controller
func New(s *mgo.Session) *Controller {
	return &Controller{s}
}

// CreateSprint e.POST("/sprints", CreateSprint)
func (sc Controller) CreateSprint(c echo.Context) error {
	s := new(models.Sprint)
	// Bind
	if err := c.Bind(s); err != nil {
		return err
	}
	// Add an Id
	s.ID = bson.NewObjectId()
	// Add creation date
	s.Created = time.Now()
	// Write the sprint to mongo
	sc.session.DB("screandb").C("sprints").Insert(s)
	// Return sprint created
	return c.JSON(http.StatusCreated, s)
}

// GetSprints e.GET("/sprints", GetSprints)
func (sc Controller) GetSprints(c echo.Context) error {
	var s []models.Sprint
	// Find Sprint By Id
	if err := sc.session.DB("screandb").C("sprints").Find(nil).Limit(10).Sort("-created").All(&s); err != nil {
		return err
	}
	// Return Sprints found
	return c.JSON(http.StatusOK, s)
}

// GetSprint e.GET("/sprints/:id", GetSprint)
func (sc Controller) GetSprint(c echo.Context) error {
	id := c.Param("id")

	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		// TODO : Error handler
		return c.String(http.StatusBadRequest, "Not a valid Id")
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	s := new(models.Sprint)
	// Find Sprint By Id
	if err := sc.session.DB("screandb").C("sprints").FindId(oid).One(&s); err != nil {
		return err
	}
	// Return Sprint found
	return c.JSON(http.StatusOK, s)
}

// UpdateSprint e.update("/sprints/:id", UpdateSprint)
func (sc Controller) UpdateSprint(c echo.Context) error {

	id := c.Param("id")

	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		// TODO : Error handler
		return c.String(http.StatusBadRequest, "Not a valid Id")
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Updated sprint
	s := new(models.Sprint)
	// Bind
	if err := c.Bind(s); err != nil {
		return err
	}

	// Update sprint
	if err := sc.session.DB("screandb").C("sprints").UpdateId(oid, s); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

// RemoveSprint e.DELETE("/sprints/:id", RemoveSprint)
func (sc Controller) RemoveSprint(c echo.Context) error {

	id := c.Param("id")

	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		// TODO : Error handler
		return c.String(http.StatusBadRequest, "Not a valid Id")
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove sprint
	if err := sc.session.DB("screandb").C("sprints").RemoveId(oid); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
