package scheduler

import "crawl/engine"

type QueueScheduler struct{
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan<-r//add request into channel
}
func (q *QueueScheduler) WorkerReady(w chan engine.Request){
	q.workerChan<-w
}

func (q *QueueScheduler) WorkerChan()chan engine.Request {
	return make(chan engine.Request)//factory
}

func (q *QueueScheduler) Run(){
		q.workerChan=make(chan  chan engine.Request)
		q.requestChan=make(chan engine.Request)
	go func(){
		var requestsQueue []engine.Request
		var workerQueue []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if(len(requestsQueue)>=1 && len(workerQueue)>=1){
				activeRequest=requestsQueue[0]
				activeWorker=workerQueue[0]
			}
			select{
				case requests:=<-q.requestChan:
					requestsQueue=append(requestsQueue,requests)
				case workers:=<-q.workerChan:
					workerQueue=append(workerQueue,workers)
				case activeWorker<-activeRequest:
					//fetch out of the queue
					requestsQueue=requestsQueue[1:]
					workerQueue=workerQueue[1:]
			}
		}
	}()
}

