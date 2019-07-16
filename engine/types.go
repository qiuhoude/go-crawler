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

// 序列化和反序列化接口
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

type ParserFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	parserFunc ParserFunc
	name       string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parserFunc(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parserFunc: p,
		name:       name,
	}
}
