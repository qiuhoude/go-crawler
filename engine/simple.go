package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (s *SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		/*
			log.Printf("Fetching  %s\n", r.Url)
			contents, err := fetcher.Fetch(r.Url)
			if err != nil {
			log.Printf("Fetch: error fetching url %s %v", r.Url, err)
			continue
			}
			parseResult := r.ParserFunc(contents)
		*/
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

//干抓取和解析工作
func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching  %s\n", r.Url)
	// 抓取
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}
	// 解析
	parseResult := r.ParserFunc(contents)
	return parseResult, nil
}
