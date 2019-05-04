package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

//用于存储Item

// 创建chan
func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d  item : %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}
		}

	}()
	return out
}

// 保存item
func save(item interface{}) (id string, err error) {
	// 关闭内网的sniff
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		//log.Println(err)
		return "", err
	}

	resp, err := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
		Index("datint_profile").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		//log.Println(err)
		return "", err
	}
	//fmt.Printf("%+v", resp) //格式化输出结构体对象的时候包含了字段名称
	return resp.Id, nil
}
