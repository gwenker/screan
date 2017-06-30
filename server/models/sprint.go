package models

import (
	"time"

	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Sprint define an iteration of development
type Sprint struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Created     time.Time     `json:"created" bson:"created"`
	Name        string        `json:"name" bson:"name"`
	StartDate   time.Time     `json:"startDate" bson:"startDate"`
	EndDate     time.Time     `json:"endDate" bson:"endDate"`
	Stream      Stream        `json:"stream" bson:"stream"`
	Capacity    float32       `json:"capacity" bson:"capacity"`
	Velocity    float32       `json:"velocity" bson:"velocity"`
	Planning    []PlannedDay  `json:"planning" bson:"planning"`
	Events      []Event       `json:"events" bson:"events"`
	UserStories []UserStory   `json:"userStories"`
}

// SprintsRepo is the a repo for sprints
type SprintsRepo interface {
	// Create a sprint
	Create(sprint Sprint) (Sprint, error)
	// FindByID the sprint by Id
	FindByID(id string) (Sprint, error)
	// FindAll get all sprints
	FindAll() ([]Sprint, error)
	// Update a sprint
	Update(id string, sprint Sprint) (Sprint, error)
	// Delete a sprint
	Delete(id string) error
}

// DefaultSprintsRepo is the repository for sprints
type DefaultSprintsRepo struct {
	coll *mgo.Collection
}

// NewSprintsRepo instantiate new SprintsRepo
func NewSprintsRepo(coll *mgo.Collection) SprintsRepo {
	return &DefaultSprintsRepo{coll: coll}
}

// Create a sprint
func (sr DefaultSprintsRepo) Create(sprint Sprint) (Sprint, error) {
	return sprint, sr.coll.Insert(sprint)
}

// FindByID the sprint by Id
func (sr DefaultSprintsRepo) FindByID(id string) (Sprint, error) {
	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		return Sprint{}, errors.New("invalid id")
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	s := Sprint{}
	// Find Sprint By Id
	err := sr.coll.FindId(oid).One(&s)
	return s, err
}

// FindAll get all sprints
func (sr DefaultSprintsRepo) FindAll() ([]Sprint, error) {
	var s []Sprint
	// Find all Sprints
	err := sr.coll.Find(bson.M{}).Limit(10).Sort("-created").All(&s)
	return s, err
}

// Update a sprint
func (sr DefaultSprintsRepo) Update(id string, sprint Sprint) (Sprint, error) {
	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		return Sprint{}, errors.New("invalid id")
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Update Sprint Bwithy Id
	err := sr.coll.UpdateId(oid, sprint)
	return sprint, err
}

// Delete a sprint
func (sr DefaultSprintsRepo) Delete(id string) error {
	// Verify id is ObjectId
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid id")
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Update Sprint Bwithy Id
	return sr.coll.RemoveId(oid)
}
