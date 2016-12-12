package sprint

import (
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gwenker/screan/server/models"
	"github.com/labstack/echo"
)

type SprintController struct {
	session *mgo.Session
}

func New(s *mgo.Session) *SprintController {
	return &SprintController{s}
}

// e.POST("/sprints", CreateSprint)
func (sc SprintController) CreateSprint(c echo.Context) error {
	s := new(models.Sprint)
	// Bind
	if err := c.Bind(s); err != nil {
		return err
	}
	// Add an Id
	s.Id = bson.NewObjectId()
	// Add creation date
	s.Created = time.Now()
	// Write the sprint to mongo
	sc.session.DB("screandb").C("sprints").Insert(s)
	// Return sprint created
	return c.JSON(http.StatusCreated, s)
}

// e.GET("/sprints", GetSprints)
func (sc SprintController) GetSprints(c echo.Context) error {
	var s []models.Sprint
	// Find Sprint By Id
	if err := sc.session.DB("screandb").C("sprints").Find(nil).Limit(10).Sort("-created").All(&s); err != nil {
		return err
	}
	// Return Sprints found
	return c.JSON(http.StatusOK, s)
}

// e.GET("/sprints/:id", GetSprint)
func (sc SprintController) GetSprint(c echo.Context) error {
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

// e.DELETE("/sprints/:id", RemoveSprint)
func (sc SprintController) RemoveSprint(c echo.Context) error {

	id := c.Param("id")

	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		// TODO : Error handler
		return c.String(http.StatusBadRequest, "Not a valid Id")
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := sc.session.DB("screandb").C("sprints").RemoveId(oid); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
