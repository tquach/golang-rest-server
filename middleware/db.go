package middleware

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

func DB(databaseUrl, databaseName string) martini.Handler {
	session, err := mgo.Dial(databaseUrl)
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		// Clone the session.
		s := session.Clone()

		// Map a reference to the database to the request context
		c.Map(s.DB(databaseName))

		// Clean up our session.
		defer s.Close()

		// Pass control to next handler
		c.Next()
	}
}
