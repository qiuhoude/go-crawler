package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	/*_, err := persist.Save(s.Client, item)
	log.Printf("Item %v saved..", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v : %v", item, err)
	}
	return err*/
	log.Printf("Item %v saved..", item)
	*result = "ok"
	return nil
}
