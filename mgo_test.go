package beemgo

import (
	"testing"
	. "gopkg.in/check.v1"
	"labix.org/v2/mgo"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&MgoSuite{})

type MgoSuite struct {
	sess *mgo.Session
}

func (s *MgoSuite) TestNew(c *C) {
	var err error
	//Actually need a live mongo server to test this
	s.sess, err = Dial("localhost")
	if err != nil {
		c.Log(err.Error())
	}
	s1 := s.sess
	s2 := Session()
	c.Assert(s2, FitsTypeOf, &mgo.Session{})
	c.Assert(s1, Not(Equals), s2)
}

func (s *MgoSuite) TestSession(c *C) {
}
