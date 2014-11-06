package handlers

import (
	"html/template"
	"net/http"
	"testing"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "gopkg.in/check.v1"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	TEST_DB          = "test"
	TEST_RESOURCE    = "notes"
	TEST_RESOURCE_ID = "52cef5dde8aa96cb99e4a68b"
)

type stubrender struct {
	status   int
	response interface{}
}

func (r *stubrender) JSON(status int, v interface{}) {
	r.status = status
	r.response = v
}
func (r *stubrender) HTML(status int, name string, v interface{}, htmlOpt ...render.HTMLOptions) {}
func (r *stubrender) Error(status int)                                                           {}
func (r *stubrender) Redirect(location string, status ...int)                                    {}
func (r *stubrender) Template() *template.Template {
	return nil
}
func (r *stubrender) Data(int, []byte) {
}
func (r *stubrender) Header() (h http.Header) {
	return
}
func (r *stubrender) Status(int) {
}
func (r *stubrender) XML(int, interface{}) {
}

type Note struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Note string        `bson:"note"`
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type HandlersTestSuite struct {
	session *mgo.Session
}

var _ = Suite(&HandlersTestSuite{})

func (s *HandlersTestSuite) SetUpSuite(c *C) {
	session, err := mgo.Dial("local-api.crowdsurge.com")
	if err != nil {
		c.Fatal("No database available.")
	}

	c.Assert(session, NotNil)
	s.session = session
}

func (s *HandlersTestSuite) TearDownSuite(c *C) {
	defer s.session.Close()

	s.session.DB(TEST_DB).C(TEST_RESOURCE).DropCollection()
}

func (s *HandlersTestSuite) TestGetResource(c *C) {
	dbSession := s.session.Clone()
	defer dbSession.Close()

	// Prime database with some test objects
	id := bson.ObjectIdHex(TEST_RESOURCE_ID)
	note := Note{Id: id, Note: "This is a note."}
	_, err := dbSession.DB(TEST_DB).C(TEST_RESOURCE).Upsert(bson.M{"_id": id}, &note)

	if err != nil {
		c.Fatalf("DB Failed %s", err.Error())
		c.FailNow()
	}

	params := make(martini.Params)
	params["id"] = TEST_RESOURCE_ID
	params["resource"] = TEST_RESOURCE

	renderer := &stubrender{}
	GetResource(params, renderer, dbSession.DB(TEST_DB))

	c.Assert(renderer.status, Equals, 200, Commentf("Should return object but got %T.", renderer.response))
}
