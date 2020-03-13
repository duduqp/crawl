package engine


type Item struct{
	Url string
	Type string
	Id string
	Payload interface{}
}


type ParserResult struct{
	Requests []Request
	Items    []Item
}


type Request struct{
	Url string
	ParserFunc func(c []byte) ParserResult
}

func NilParser([]byte) ParserResult{
	return ParserResult{}
}