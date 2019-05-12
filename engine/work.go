package engine

import (
	"crawler/fetcher"
	"log"
)

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
