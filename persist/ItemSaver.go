package persist

import (
	"context"
	"crawl/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item,error){
	client,err:=elastic.NewClient(elastic.SetSniff(false))
	if err!=nil{
		//todo
		return nil,err
	}

	out:=make(chan engine.Item)
	go func(){
		itemcount:=0
		for{
			item:=<-out//绑定
			log.Printf("Item Saver Got Item: #%d: %v",itemcount,item)
			itemcount++

			_,err:=save(client,item,index)
			if err!=nil{
				//give up
				log.Printf("Item Saver:error"+"saving Item :%v  with %v",item,err)
			}
		}
	}()
	return out,nil //传出去给别人放item到通道中
}

//return saved id
func save(client * elastic.Client,item engine.Item,index string)(id string ,err error) {

	//Try start up elastic search
	//docker go client
	//client,err:=elastic.NewClient(elastic.SetSniff(false))
	//if err!=nil{
	//	//todo
	//	return "",err
	//}

	if item.Type==""{
		return "",errors.New("must have type")
	}

	indexService:=client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id!=""{
		indexService.Id(item.Id)
	}



	resp,err:=indexService.
		Do(context.Background())
	if err!=nil{
		return "",err
	}

	return resp.Id,nil
}


