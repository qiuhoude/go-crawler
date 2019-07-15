package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist/client"
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
	"fmt"
	"log"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        url,
	//	ParserFunc: parser.ParseCityList,
	//})
	//itemSaver, err := persist.ItemSaver()
	itemSaver, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
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
