package parser

import (
	"crawl/engine"
	"crawl/model"
	"regexp"
	"strconv"
)



var pattern_profile_Name= regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)

var pattern_profile_Name_More= regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">`)

func ParseUser(content []byte) engine.ParserResult{
	matches:=pattern_profile_Name.FindAllSubmatch(content,-1)

	res:=engine.ParserResult{}
	for _,m:=range matches{
		name:=string(m[2])
		url:=string(m[1])
		//res.Items=append(res.Items,"User:"+string(m[2]))
		res.Requests =append(res.Requests,engine.Request{
			Url:       url,
			ParserFunc:func(c []byte) engine.ParserResult{
				return ParseProfile(c,name,url)
			},
		})
	}

	matches=pattern_profile_Name_More.FindAllSubmatch(content,-1)
	for _,m:=range matches{
		res.Requests =append(res.Requests,engine.Request{
			Url:       string(m[1]),
			ParserFunc:ParseUser,
		})
	}
	return res
}
const City_re =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseInfoList(content []byte) engine.ParserResult{
	re:=regexp.MustCompile(City_re)
	matches:=re.FindAllSubmatch(content,-1)

	result:=engine.ParserResult{}
	for _,m := range matches{
		//result.Items=append(result.Items,"City: "+string(m[2]))
		result.Requests =append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseUser,
			})
	}
	return result
}



//  Name string
//	Gender string
//	Age int
//	Height int
//	Income string
//	Marriage string
//	Hometown string
//性别：男士	居住地：西藏拉萨
//年龄：52	月   薪：12001-20000元
//婚况：离异	身   高：176
var pattern_profile_Age = regexp.MustCompile(`"age":([0-9]+),`)
var pattern_profile_Marriage= regexp.MustCompile(`"basicInfo":\["([^"]+)",`)
var pattern_profile_Income= regexp.MustCompile(`"月收入:([^"]+)"`)
var pattern_profile_Gender= regexp.MustCompile(`"genderString":"([^"]+)"`)
var pattern_profile_Height= regexp.MustCompile(`"([0-9]+)cm"`)
var pattern_profile_Hometown= regexp.MustCompile(`"籍贯:([^"]+)"`)
var idUrlRe=regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
//url末尾就是id
func ParseProfile(content []byte,name string,url string) engine.ParserResult{
	//re:=regexp.MustCompile(pattern_profile)
	//match:= pattern_profile_Age.FindSubmatch(content)
	//if match!=nil{
	//	age,err :=strconv.Atoi(extractCharacteristicString(content, pattern_profile_Age))
	//	if err!=nil{
	//		//user age is age
	//		profile.Age=age
	//	}
	//}
	profile:=model.Profile{}
	match:=extractCharacteristicString (content,pattern_profile_Age)
	if match!=""{
			profile.Age,_=strconv.Atoi(match)
	}

	match=extractCharacteristicString (content,pattern_profile_Height)
	if match!=""{
			profile.Height=match
	}
	match=extractCharacteristicString (content,pattern_profile_Gender)
	if match!=""{
			profile.Gender=match
		}

	match=extractCharacteristicString (content,pattern_profile_Income)
	if match!=""{
		profile.Income=(match)
	}

	match=extractCharacteristicString (content,pattern_profile_Hometown)
	if match!=""{
		profile.Hometown=match
	}

	match=extractCharacteristicString (content,pattern_profile_Marriage)
	if match!=""{
		profile.Marriage=match
	}

	//special case for name
	//match=extractCharacteristicString (content,pattern_profile_Name)
	//if match!=""{
	//	profile.Name=match
	//}
	profile.Name=name


	//nest decode url
	var id string = extractCharacteristicString([]byte(url),idUrlRe)
	//can also extractcharacteristic like gender ... income and set into model
	result:=engine.ParserResult{
		Items:[]engine.Item{
			{
				Url: url,
				Type:    "Lover",
				Id:      id,
				Payload: profile,
			},
		},
	}//can have more requests
	return result
	}

func extractCharacteristicString (content []byte,re * regexp.Regexp) string{
	match:=re.FindSubmatch(content)
	if len(match) >=2{
		return string(match[1])
	}
	return ""
}



