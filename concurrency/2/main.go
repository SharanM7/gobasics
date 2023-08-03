package main

import (
	"log"
	"sync"
	"sync/atomic"
)

type Server struct {
	Count int64
	Mu    sync.Mutex
}

func main() {
	s := &Server{}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go s.addWithMutex(i, &wg)
		go s.addWithAtomocity(i, &wg)
	}

	wg.Wait()
	log.Printf("value : %+v\n", s)
}

func (s *Server) addWithMutex(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	s.Mu.Lock()
	s.Count += 1
	s.Mu.Unlock()
}

func (s *Server) addWithAtomocity(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&s.Count, int64(1))
}
