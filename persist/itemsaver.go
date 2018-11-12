package persist

import (
	"context"
	"errors"
	"log"

	"github.com/olivere/elastic"
	"github.com/researchlab/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {
	if item.Type == "" {
		return errors.New("Must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v\n", resp)
	return nil
}
