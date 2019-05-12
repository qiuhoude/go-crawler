package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
	"log"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        url,
	//	ParserFunc: parser.ParseCityList,
	//})
	itemSaver, err := persist.ItemSaver()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
	}
	concurrentEngine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList4Json,
	})
}
