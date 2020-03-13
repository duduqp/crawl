package engine

import (
	"crawl/fetcher"
	"fmt"
	"log"
)
type SingleEngine struct{}
//vary parameter
func( s *SingleEngine) Run(seeds ...Request){
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}

	//while loop
	for len(requests) > 0{
		r:=requests[0]
		requests=requests[1:]//fetch first request


		//fetching url
		log.Printf("Fetching %s",r.Url)
		parserResult,err:=worker(r)
		if err!=nil{
			fmt.Println(err)
			continue //next request from list
		}

		requests = append(requests,parserResult.Requests...)

		for _,item:=range parserResult.Items{
		log.Printf("Got Item : %s",item)
		}
	}
}


func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching url: %v", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}