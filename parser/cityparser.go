package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	cityRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\w]+/[0-9]+)"[^>]*>[\s]*下一页[\s]*</a>`)
)

func ParseCity(contents []byte) engine.ParseResult {

	all := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		name := string(c[2])
		result.Items = append(result.Items, name) // 用户名称
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	return result
}
