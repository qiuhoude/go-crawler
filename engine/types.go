package engine

//解析结果
type ParseResult struct {
	Requests []Request
	Items    [] interface{}
}

type Request struct {
	Url        string                   //解析出来的URL
	ParserFunc func([]byte) ParseResult //处理这个URL所需要的函数
}

// 创建空parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
