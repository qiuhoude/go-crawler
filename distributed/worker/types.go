package worker

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/parser"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	Name string //functionName 序列化后的解析器函数
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

// 获取解析Parser函数
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil

	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil

	case config.NilParser:
		return engine.NilParser{}, nil

	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil

		} else {
			return nil, fmt.Errorf("invalid arg : %v", p.Args)
		}

	default:
		return nil, errors.New("unknown parser name")
	}
}

func SerializeParseResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error desrializing request : %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
