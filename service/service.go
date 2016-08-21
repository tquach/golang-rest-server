package service

import "gopkg.in/mgo.v2"

// RepositoryService declares an interface for data persistence operations.
type RepositoryService interface {
	Insert(resource string, data interface{}) error
	// Update(resource string, data interface{}) error
	// Delete(resource string, id string) error
	// Find(resources string, id string) (interface{}, error)
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
