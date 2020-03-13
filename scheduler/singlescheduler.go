package scheduler

import "crawl/engine"

type SingleScheduler struct {
	workerchan chan engine.Request
}

func (s *SingleScheduler) WorkerChan() chan engine.Request {
	return s.workerchan
}

func (s *SingleScheduler) WorkerReady(chan engine.Request) {
	panic("implement not")
}

func (s *SingleScheduler) Run() {
	s.workerchan=make(chan engine.Request)
}

func (s  *SingleScheduler) Submit(r engine.Request){
	go func(){s.workerchan<-r}()//dead lock
}

