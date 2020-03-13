package persist

import (
	"context"
	"crawl/engine"
	"crawl/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {

	client,err:=elastic.NewClient(elastic.SetSniff(false))

	profile:=engine.Item{
		Url:     "https://album.zhenai.com/u/1488229813",
		Type:    "Lover",
		Id:      "1488229813",
		Payload: model.Profile{
		Name:     "dudu",
		Gender:   "male",
		Age:      21,
		Height:   "172",
		Income:   "100000",
		Marriage: "no",
		Hometown: "ZheJiang",
	},
	}
	//save testcase
	id,err:=save(client,"dat_profile_test",profile)
	if err!=nil{
		panic(err)
	}

	resp,err:=client.Get().Index("dat_profile").Type(profile.Type).Id(profile.Id).Do(context.Background())

	if err!=nil{
		panic(err)
	}

	t.Logf("%s",resp.Source)

	//fetch 
	var actual engine.Item
	err=json.Unmarshal(*resp.Source,&actual)
	//把嵌套内层json单独解码成json 再注入进actual外层对象
	actualprofile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload=actualprofile
	if err!=nil{
		panic(err)
	}

	if actual!=profile{
		t.Log("wrong match")
	}

}
