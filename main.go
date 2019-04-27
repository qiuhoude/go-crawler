package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        url,
	//	ParserFunc: parser.ParseCityList,
	//})
	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	concurrentEngine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
