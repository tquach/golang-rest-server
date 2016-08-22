package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/tquach/golang-rest-server/logger"
	"github.com/tquach/golang-rest-server/render"
	"github.com/tquach/golang-rest-server/service"
)

// Controller defines functions to handle web requests.
type Controller struct {
	repo   service.RepositoryService
	logger logger.Logger
}

// FindResource will retrieve a specific resource by an id parameter.
func (c Controller) FindResource(w http.ResponseWriter, req *http.Request) {
	resource := req.URL.Query().Get(":resource")
	id := req.URL.Query().Get(":id")
	c.logger.Debugf("find resource %s/%s", resource, id)

	v, err := c.repo.Find(resource, id)
	if err != nil {
		c.logger.Errorf("could not retrieve resource %s", err)
		render.JSONError(err, http.StatusNotFound, w)
		return
	}

	render.JSON(v, http.StatusOK, w)
}

// SaveResource will parse the request body as a new document and save to the appropriate collection.
func (c Controller) SaveResource(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{}
	resource := req.URL.Query().Get(":resource")
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		render.JSONError(err, http.StatusBadRequest, w)
		return
	}

	if err := c.repo.Insert(resource, data); err != nil {
		render.JSONError(err, http.StatusBadRequest, w)
		return
	}

	render.JSON(data, http.StatusOK, w)
}

// New constructs a new instance of the controller.
func New(repo service.RepositoryService) Controller {
	return Controller{
		repo:   repo,
		logger: logger.New("ctrl"),
	}
}
