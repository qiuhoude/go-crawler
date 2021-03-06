package client

import (
	"crawler/distributed/rpcsupport"
	"crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			//log.Printf(" 发送保存数据 已发%d个 :\n", itemCount)
			itemCount++

			//调用Rpc 来保存item

			result := ""
			err = client.Call("ItemSaveService.Save", item, &result)

			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}

		}
	}()
	return out, nil

}
