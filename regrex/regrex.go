package main

const text_a string = "My QQ name is Iris-du"
const text_b string = `
this is first line
this is for match irisdu@gmail.com
this also match duduqp@qq.com
`
//func main(){
//
//	//if must be correct then you can use MustComplie
//	p,err:=regexp.Compile("Iris-du")
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//
//	match:=p.Find([]byte(text_a))
//	fmt.Printf("%s\n",match)
//	//match=p.FindString(text)
//	//p=regexp.MustCompile(`([a-zA-Z0-9])+@.+\..+`)//\.. is escape .
//	p=regexp.MustCompile(`([a-zA-Z0-9]+)@.+\..+`)
//	match_b:=p.FindAllStringSubmatch(text_b,-1)//end index -1
//	for _,ss := range match_b{
//		fmt.Println(ss)
//	}
//	return
//}
