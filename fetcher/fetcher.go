package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//must big case if should exposed to other package
func Fetch(url string) ([]byte ,error) {

	//Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36
	//res, err := http.Get(url)
	//
	//if err != nil {
	//	//panic(err)
	//	return nil,err
	//}
	//defer res.Body.Close()
	//403
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	//judge status
	if res.StatusCode != http.StatusOK {
		//if web content is gbk
		//utf8rd:=transform.NewReader(res.Body,simplifiedchinese.GBK.NewEncoder())
		//content,err:=ioutil.ReadAll(utf8rd)

		//auto detect encoding #

		fmt.Println("Error: status code",res.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d",res.StatusCode)
	}
	content, err := ioutil.ReadAll(res.Body)
	//printInfoList(content)
	if err != nil {
		//panic(err)
		return nil,err
	}
	//fmt.Printf("%s\n",content)
	return content,err


}