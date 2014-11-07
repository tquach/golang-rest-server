package handlers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-martini/martini"
	. "gopkg.in/check.v1"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	TEST_DB          = "test"
	TEST_RESOURCE    = "notes"
	TEST_RESOURCE_ID = "52cef5dde8aa96cb99e4a68b"
)

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
	session, err := mgo.Dial("localhost")
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

func (s *HandlersTestSuite) TestSaveResource(c *C) {
	dbSession := s.session.Clone()
	defer dbSession.Close()

	params := make(martini.Params)
	params["id"] = TEST_RESOURCE_ID
	params["resource"] = TEST_RESOURCE

	renderer := &stubrender{}
	reader := strings.NewReader("{\"note\": \"This is a new note\"}")
	req, _ := http.NewRequest("POST", "/notes", reader)
	SaveResource(params, req, renderer, dbSession.DB(TEST_DB))

	cnt, _ := dbSession.DB(TEST_DB).C(TEST_RESOURCE).Find(bson.M{"note": "This is a new note"}).Count()
	c.Assert(cnt, Equals, 1)
}
