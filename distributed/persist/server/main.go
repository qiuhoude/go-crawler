package main

import (
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		*port = 7788
	}
	//如果发生错误，Fatal()会强制退出。。
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), "datint_profile"))
}
func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
}
