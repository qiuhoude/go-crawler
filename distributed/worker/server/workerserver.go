package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		*port = config.WorkerPort0
	}

	workerPort := fmt.Sprintf(":%d", *port)
	//发布worker服务
	log.Fatal(rpcsupport.ServeRpc(workerPort, worker.CrawlService{}))
}
