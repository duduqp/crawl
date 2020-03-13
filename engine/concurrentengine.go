package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}


type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan()chan Request
	Run()
}


type ReadyNotifier interface {
	WorkerReady(chan Request)
}
func (c * ConcurrentEngine) Run(seeds ...Request){


	//in:=make(chan Request)
	out:=make(chan ParserResult)
	c.Scheduler.Run()
	for i:=0 ; i<c.WorkerCount;i++{
		createWorker(c.Scheduler.WorkerChan(),out,c.Scheduler)
	}//concurrent by create goroutine but can not handler subroutine
	//can cache the requests!both cache the worker who is idle!
	//who router the request to worker? scheduler!
	for _,seed :=range seeds{
		c.Scheduler.Submit(seed)
	}
	ItemCount:=0
	for{
		result:=<-out
		for _,item:=range result.Items{
			ItemCount++
			//log.Printf("Got #%d Item from channel:%v",ItemCount,item)
			go func(){c.ItemChan<-item}()
		}

		for _,request:= range result.Requests{
			c.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request,out chan ParserResult,r ReadyNotifier){
	//in:=make(chan Request)
	go func(){
		for{
			//signal scheduler i am ready
			r.WorkerReady(in)
			request:=<-in
			result,err:=worker(request)
			if err!=nil{
				continue
			}
			out<-result//dead lock
		}
	}()
}