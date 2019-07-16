package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist/client"
	"crawler/distributed/rpcsupport"
	client2 "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

func printItem() (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		for item := range out {
			fmt.Println(item)
		}
	}()
	return out, nil
}

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHost    = flag.String("worker_hosts", "", "worker hosts (comma separated)") //逗号分隔
)

func main() {
	flag.Parse()
	url := "http://www.zhenai.com/zhenghun"

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        url,
	//	ParserFunc: parser.ParseCityList,
	//})
	//itemSaver, err := persist.ItemSaver()

	// 存储的rpc
	if len(*itemSaverHost) == 0 {
		*itemSaverHost = fmt.Sprintf(":%d", config.ItemSaverPort)
	}
	itemSaver, err := client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	// 请求和解析rpc
	pool := createClientPool(strings.Split(*workerHost, ","))
	processor, err := client2.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
		RequestProcessor: processor,
	}
	concurrentEngine.Run(engine.Request{
		Url: url,
		//Parser: parser.ParseCityList4Json ,
		Parser: engine.NewFuncParser(parser.ParseCityList4Json, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	// 创建client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s : %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
