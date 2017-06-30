package db

import (
	"net/http"

	"github.com/gwenker/screan/server/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"

	"strings"
	"time"
)

const mongoTimeout = 10 * time.Second

// Session stores mongo session
var session *mgo.Session

//Client is the entrypoint of Docktor API
type Client interface {
	Sprints() models.SprintsRepo
	Close()
}

// DB is the implementation structure to call mongo
type DB struct {
	session *mgo.Session
	sprints models.SprintsRepo
}

// Close the connexion to mongo
func (db *DB) Close() {
	db.session.Close()
}

// Sprints is the entrypoint for sprints model
func (db *DB) Sprints() models.SprintsRepo {
	return db.sprints
}

// Connect connects to mongodb
func Connect(uri string) {
	s, _ := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   strings.Split(uri, ","),
		Timeout: mongoTimeout,
	})
	s.SetSafe(&mgo.Safe{})
	session = s
}

// Get the connexion to docktor API
func Get(username, password string) (Client, error) {
	s := session.Clone()
	s.SetSafe(&mgo.Safe{})
	database := s.DB("screandb")
	if username != "" && password != "" {
		err := database.Login(username, password)
		if err != nil {
			return nil, err
		}
	}

	return &DB{
		sprints: models.NewSprintsRepo(database.C("sprints")),
		session: s,
	}, nil
}

// Close the connexion to mongo
func Close() {
	if session != nil {
		session.Close()
	}
}

// Middleware to add db in echo.context
func Middleware(username, password string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db, err := Get(username, password)
			if err != nil {
				return c.JSON(http.StatusServiceUnavailable, map[string]string{
					"message": err.Error(),
				})
			}
			defer db.Close()
			c.Set("db", db)
			return next(c)
		}
	}
}
