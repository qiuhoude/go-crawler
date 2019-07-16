package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	// 发布服务
	host := fmt.Sprintf(":%d", config.WorkerPort0)
	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(3 * time.Second)
	// 客户端调用
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1214814888",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "林YY",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
