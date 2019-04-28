package persist

import "log"

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
		}

	}()
	return out
}
