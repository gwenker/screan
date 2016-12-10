package sprint

import (
	"net/http"

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
	// Write the sprint to mongo
	sc.session.DB("screandb").C("sprints").Insert(s)
	// Return sprint created
	return c.JSON(http.StatusCreated, s)
}

// e.GET("/sprints", GetSprints)
func (sc SprintController) GetSprints(c echo.Context) error {
	return c.String(http.StatusOK, "sprints")
}

// e.GET("/sprints/:id", GetSprint)
func (sc SprintController) GetSprint(c echo.Context) error {
	id := c.Param("id")

	// Verify id is ObjectId, otherwise bail
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
