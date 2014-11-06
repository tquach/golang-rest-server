// Package `handlers` provides single retrieval or list of resources based on the URI path.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func SaveResource(p martini.Params, req *http.Request, r render.Render, db *mgo.Database) {
	var data map[string]interface{}

	resource := p["resource"]
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&data)
	if err == nil {
		db.C(resource).Insert(data)
		r.JSON(201, &data)
	} else {
		r.JSON(400, map[string]string{"error": err.Error()})
	}
	return
}

// Retrieve an instance of the resource given the id. Assumes the resource name matches the collection name
func GetResource(p martini.Params, r render.Render, db *mgo.Database) {
	resource := p["resource"]
	id := p["id"]

	// TODO use reflection
	var result *interface{}
	if !bson.IsObjectIdHex(id) {
		r.JSON(400, map[string]string{"error": "Invalid id"})
		return
	}

	err := db.C(resource).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		var status int
		if err == mgo.ErrNotFound {
			status = 404
		} else {
			status = 500
		}

		r.JSON(status, map[string]string{"error": err.Error()})
		return
	}
	r.JSON(200, result)
}

// Retrieve a collection of resources based on the params
func ListResources(p martini.Params, r render.Render, db *mgo.Database) {
	resource := p["resource"]
	var results []interface{}
	err := db.C(resource).Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	r.JSON(200, &results)
}
