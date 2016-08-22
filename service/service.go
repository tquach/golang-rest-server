package service

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RepositoryService declares an interface for data persistence operations.
type RepositoryService interface {
	Insert(resource string, data interface{}) error
	// Update(resource string, data interface{}) error
	// Delete(resource string, id string) error
	Find(resource string, id string) (interface{}, error)
}

// MongoRepository implements a RepositoryService with a Mongo DB backing store.
type MongoRepository struct {
	db *mgo.Database
}

// Insert inserts a new document in the database for the given resource collection.
func (m MongoRepository) Insert(resource string, data interface{}) error {
	c := m.db.C(resource)
	return c.Insert(&data)
}

// Find retrieves the document with the given id.
func (m MongoRepository) Find(resource string, id string) (interface{}, error) {
	c := m.db.C(resource)

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("Invalid bson ID")
	}

	result := map[string]interface{}{}
	if err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// NewMongoRepository constructs a new instance of MongoRepository
func NewMongoRepository(url string, dbName string) (MongoRepository, error) {
	session, err := mgo.Dial(url)

	// Ignore errors from write operations.
	session.SetSafe(nil)

	if err != nil {
		return MongoRepository{}, err
	}

	m := MongoRepository{
		db: session.DB(dbName),
	}

	return m, nil
}
