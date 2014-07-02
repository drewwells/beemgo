package beemgo

import (
	"log"

	"labix.org/v2/mgo"
)

var max = 10

// Service manages the pool of mongo connections.
type Service struct {
	spark *mgo.Session
	queue chan int
	Url   string
	Open  int
}

// New bootstraps the Mongo pool making it possible to open
// sessions.
func (s *Service) New() {
	var err error
	s.queue = make(chan int, max)
	for i := 0; i < max; i = i + 1 {
		s.queue <- i
	}
	s.Open = 0
	s.spark, err = mgo.Dial(s.Url)
	if err != nil {
		log.Fatal(err)
	}
}

// Session attempts to create a session in the pool.
func (s *Service) Session() *mgo.Session {
	<-s.queue
	s.Open++
	sess := s.spark.Copy()
	return sess
}

// Close an open session to free up space in the pool.
func (s *Service) Close(m *mgo.Session) {
	m.Close()
	s.queue <- 1
	s.Open--
}
