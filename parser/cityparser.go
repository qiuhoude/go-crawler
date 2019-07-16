package parser

import (
	"crawler/distributed/config"
	"crawler/engine"
	"regexp"
)

var (
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	//cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\w]+/[0-9]+)"[^>]*>[\s]*下一页[\s]*</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\w]+)`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {

	all := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		url := string(c[1])
		name := string(c[2])
		//result.Items = append(result.Items, name) // 用户名称

		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			//ParserFunc: func(bytes []byte) engine.ParseResult {
			//	return ParseProfile(bytes, url, name)
			//},
			Parser: NewProfileParser(name),
		})
	}
	// 下一页的url添加到Requests中
	nextUrls := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, c := range nextUrls {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			//ParserFunc:ParseCity,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
