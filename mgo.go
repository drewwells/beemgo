package beemgo

import (
	"log"

	"labix.org/v2/mgo"
)

type Mongo struct {
	Session *mgo.Session
}

var (
	m = Mongo{}
)

func Dial(url string) (*mgo.Session, error) {
	session, err := mgo.Dial(url)
	session.SetMode(mgo.Monotonic, true)
	m.Session = session
	return session, err
}

//Return a copy of the original mongo connection parameters
func Session() *mgo.Session {
	if m.Session == nil {
		log.Fatal("Session doesn't exist")
	}
	return m.Session.New()
}

func Close(s *mgo.Session) {
	defer func() {
		s.Close()
	}()
}
