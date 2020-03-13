package main

import (
	"crawl/engine"
	"crawl/parser"
	"crawl/persist"
	"crawl/scheduler"
	"fmt"
	"regexp"
)

func main(){
	itemChan,err:=persist.ItemSaver("dat_profile")
	if err!=nil{
		panic(err)
	}
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueueScheduler{},
		WorkerCount:1,
		ItemChan:itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseUser,
	})
	//locate html elem
	//css selector $('#1>1a>1aa...')
	//xpath
	//regrex *

	return
}


func printInfoList(content []byte){
	//match the city info href
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/zhengzhou" data-v-5fa74e39>郑州</a>
	pattern:=regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	match:=pattern.FindAll(content,-1)

	for _,ss:=range match{
		fmt.Printf("%s\n",ss)
	}
	fmt.Println(len(match))
	pattern=regexp.MustCompile(`<a href=("http://www.zhenai.com/zhenghun/[0-9a-z]+")[^>]*>([^<]+)</a>`)
	match_detail:=pattern.FindAllSubmatch(content,-1)
	//match_detail [][]string [][][]byte

	for _,m:=range match_detail{
		//for _,subm:=range m{
		//	fmt.Printf("%s",subm)
		//}
			fmt.Printf("whole  :  %s  |city  :  %s|url   :   %s|",m[0],m[2],m[1])

		fmt.Println()
	}


}

//func detectEncoding(r io.Reader) encoding.Encoding{
//	bytes,err:=bufio.NewReader(r).Peek(1024)
//	if err !=nil{
//		panic(err)
//	}
//	charset.DetermineEncoding()
//
//}