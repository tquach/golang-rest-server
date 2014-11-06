package main

import (
	"flag"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/tquach/golang-rest-server/handlers"
	"github.com/tquach/golang-rest-server/middleware"
)

// DATABASE_URL is of the form:
//    [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
var (
	databaseUrl  = flag.String("databaseUrl", "localhost", "url of database, eg. mongodb://localhost:27017")
	databaseName = flag.String("databaseName", "appDb", "database name, eg. myDatabase")
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
	s.Post("/:resource", handlers.SaveResource)
}

func AppServer() *Server {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.Use(middleware.Nsq())
	m.Use(middleware.DB(*databaseUrl, *databaseName))

	m.Action(r.Handle)
	return &Server{m, r}
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
