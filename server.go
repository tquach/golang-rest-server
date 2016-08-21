package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/pat"
	"github.com/tquach/golang-rest-server/controller"
	"github.com/tquach/golang-rest-server/service"
)

// DATABASE_URL is of the form:
//    [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
var (
	hostname     = flag.String("hostname", "localhost:9000", "hostname to bind to")
	databaseURL  = flag.String("databaseURL", "localhost", "url of database, eg. mongodb://localhost:27017")
	databaseName = flag.String("databaseName", "appDb", "database name, eg. myDatabase")
)

// Opts captures server options.
type Opts struct {
	url string
}

// Middleware is a wrapper for HTTP handler functions.
type Middleware func(h http.Handler) http.Handler

// Server defines the attributes in a web application server.
type Server struct {
	*pat.Router
	middleware []Middleware
}

// Query contains properties for defining a database query.
type Query struct {
	MaxResults int `form:"max_results"`
	Page       int `form:"page"`
	Offset     int `form:"offset"`
}

// Use will append any middleware handlers.
func (s *Server) Use(m Middleware) {
	s.middleware = append(s.middleware, m)
}

// Start will chain all middleware and start up the server.
func (s *Server) Start(bindURL string) error {
	log.Printf("Starting server on %s...\n", bindURL)
	return http.ListenAndServe(bindURL, s.Router)
}

// Ping returns a ping to the server.
func Ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}

// NewServer creates a new instance of the Server.
func NewServer() *Server {
	r := pat.New()
	r.Get("/status", Ping)

	s := &Server{
		Router: r,
	}

	// Add common middleware components
	s.Use(handlers.CORS())
	return s
}

func main() {
	// Parse command line arguments
	flag.Parse()

	// Init the app server
	s := NewServer()

	repo, err := service.NewMongoRepository(*databaseURL, *databaseName)
	if err != nil {
		log.Fatal(err)
	}

	ctrl := controller.New(repo)

	// Add some routes
	s.Post("/:resource", ctrl.SaveResource)

	// Start up the server
	log.Fatal(s.Start(*hostname))
}
