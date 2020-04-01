package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error)  {
	out := make(chan engine.Item)

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out

			if item.Payload == nil {
				continue
			}

			profile, ok := item.Payload.(model.Profile)
			if !ok {
				continue
			}
			if profile.Age == "" {
				continue
			}

			log.Printf("Item Saver:got item "+"#%d: %v", itemCount, item)
			itemCount ++

			err := Save(index, client, item)
			if err != nil {
				log.Printf("Item Saver: error " + "saving item %v:%v", item, err)
			}
		}
	}()

	return out, nil
}


func Save(index string, client *elastic.Client , item engine.Item) ( err error) {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.BodyJson(item).
		Do(context.Background())

	if err != nil {
		return  err
	}

	return nil
}