package engine

//解析结果
type ParseResult struct {
	Requests []Request
	//Items    [] interface{}
	Items []Item
}

type Request struct {
	Url string //解析出来的URL
	//ParserFunc func([]byte) ParseResult //处理这个URL所需要的函数
	Parser Parser
}

// 创建空parser
//func NilParser([]byte) ParseResult {
//	return ParseResult{}
//}

type Item struct {
	Url     string
	Index   string //存储到ElasticSearch时的type,es6之后不支持type了，所以改成index
	Id      string //用户Id
	PayLoad interface{}
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//空parser
type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}
