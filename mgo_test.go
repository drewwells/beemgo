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
	ch := svc.Session()
	s1 := <-ch
	c.Assert(svc.Open, Equals, 1)
	svc.Close(s1)
	qSize := 0
	for _, v := range svc.pool {
		if v != nil {
			qSize++
		}
	}
	c.Assert(qSize, Equals, 1)
	c.Assert(svc.Open, Equals, 0)
	// Actually need a live mongo server to test this
	// s.sess, err = Dial("localhost")
	// if err != nil {
	// 	c.Log(err.Error())
	// }
	// s1 := s.sess
	// s2 := Session()
	// c.Assert(s2, FitsTypeOf, &mgo.Session{})
	// c.Assert(s1, Not(Equals), s2)

}
