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
	svc := Service{
		Url: "localhost",
	}
	svc.New()
	s1 := svc.Session()
	c.Assert(svc.Open, Equals, 1)
	s2 := svc.Session()
	c.Assert(svc.Open, Equals, 2)
	svc.Close(s1)
	c.Assert(svc.Open, Equals, 1)
	svc.Close(s2)
	var ref *mgo.Session
	delay := false
	for i := 0; i < 10; i = i + 1 {
		ref = svc.Session()
	}
	go func() {
		//This will hang
		svc.Session()
		c.Assert(delay, Equals, true)
	}()
	svc.Close(ref)
	delay = true
}
