package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/gorilla/handlers"
	"github.com/tquach/golang-rest-server/ctrl"
	"github.com/tquach/golang-rest-server/logger"
	"github.com/tquach/golang-rest-server/service"
)

// DATABASE_URL is of the form:
//    [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
var (
	hostname     = flag.String("hostname", "localhost:9000", "hostname to bind to")
	databaseURL  = flag.String("databaseURL", "localhost", "url of database, eg. mongodb://localhost:27017")
	databaseName = flag.String("databaseName", "appDb", "database name, eg. myDatabase")
)

// Middleware is a wrapper for HTTP handler functions.
type Middleware func(h http.Handler) http.Handler

// Server defines the attributes in a web application server.
type Server struct {
	mux        *pat.PatternServeMux
	logger     logger.Logger
	middleware []Middleware
}

// Use will append any middleware handlers.
func (s *Server) Use(m Middleware) {
	s.middleware = append(s.middleware, m)
}

// AddRoute will map a handler to an endpoint and HTTP method. All HTTP methods supported.
func (s *Server) AddRoute(method string, path string, hdlr http.HandlerFunc) {
	s.mux.Add(method, path, hdlr)
}

// Get adds a GET handler on the given path.
func (s *Server) Get(path string, hdlr http.HandlerFunc) {
	s.AddRoute("GET", path, hdlr)
}

// Post adds a POST handler on the given path.
func (s *Server) Post(path string, hdlr http.HandlerFunc) {
	s.AddRoute("POST", path, hdlr)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.logger.Infof("Serving request")
	s.mux.ServeHTTP(w, req)
}

// Start will chain all middleware and start up the server.
func (s *Server) Start(bindURL string) error {
	s.logger.Infof("Starting server on %s...\n", bindURL)
	return http.ListenAndServe(bindURL, s)
}

// NewServer creates a new instance of the Server.
func NewServer() *Server {
	s := &Server{
		mux:    pat.New(),
		logger: logger.New("app-server"),
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

	c := ctrl.New(repo)

	// Add some routes
	s.Get("/:resource/:id", c.FindResource)
	s.Post("/:resource", c.SaveResource)

	// Start up the server
	log.Fatal(s.Start(*hostname))
}
