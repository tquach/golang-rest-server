package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tquach/golang-rest-server/service"
)

// Controller defines functions to handle web requests.
type Controller struct {
	repo service.RepositoryService
}

// SaveResource will parse the request body as a new document and save to the appropriate collection.
func (c Controller) SaveResource(w http.ResponseWriter, req *http.Request) {
	var data map[string]interface{}

	resource := req.URL.Query().Get(":resource")
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {

	}

	if err := c.repo.Insert(resource, data); err != nil {

	}
}

// New constructs a new instance of the controller.
func New(repo service.RepositoryService) Controller {
	return Controller{
		repo: repo,
	}
}
