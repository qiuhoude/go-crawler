package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

const EsUrl = "http://localhost:9200"

//用于存储Item

// 创建chan
func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	// 关闭内网的sniff
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(EsUrl))
	if err != nil {
		return nil, err
	}

	// 创建一个协程进行数据保存
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d  item : %v", itemCount, item)
			itemCount++

			//存储到es中
			_, err := Save(client, item)
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}

		}

	}()
	return out, nil
}

// 保存item
func Save(client *elastic.Client, item engine.Item) (id string, err error) {
	if item.Index == "" {
		return "", errors.New("must supply Index ...")
	}

	indexService := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
					Index(item.Index).
					BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())

	if err != nil {
		//log.Println(err)
		return "", err
	}
	fmt.Printf("%+v", resp) //格式化输出结构体对象的时候包含了字段名称
	return resp.Id, nil
}
