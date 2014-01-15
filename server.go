package main

import (
	"flag"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/tantastik/golang-rest-demo/handlers"
	"labix.org/v2/mgo"
)

// DATABASE_URL is of the form:
//    [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
var (
	databaseUrl  = flag.String("databaseUrl", "localhost", "url of database, eg. mongodb://localhost:27017")
	databaseName = flag.String("databaseName", "test", "database name, eg. myDatabase")
)

type Server struct {
	*martini.Martini
	martini.Router
}

// Example of a basic document binding
type Query struct {
	MaxResults int `form:"max_results"`
	Page       int `form:"page"`
	Offset     int `form:"offset"`
}

// Adds routes to the server
func (s *Server) AddRoutes() {
	// Bind to a basic form for passing in additional query parameters
	s.Get("/:resource/:id", binding.Bind(Query{}), handlers.GetResource)
	s.Get("/:resource", binding.Bind(Query{}), handlers.ListResources)
	s.Post("/:resource/", handlers.SaveResource)
}

func AppServer() *Server {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.Use(DB())

	m.Action(r.Handle)
	return &Server{m, r}
}

func DB() martini.Handler {
	session, err := mgo.Dial(*databaseUrl)
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		// Clone the session.
		s := session.Clone()

		// Map a reference to the database to the request context
		c.Map(s.DB(*databaseName))

		// Clean up our session.
		defer s.Close()

		// Pass control to next handler
		c.Next()
	}
}

func main() {
	// Parse command line arguments
	flag.Parse()

	// Init the app server
	s := AppServer()

	// Add some routes
	s.AddRoutes()

	// Run all the things
	s.Run()
}
