package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/RideForJoy/parse-pc-files/internal/product"
)

const (
	projectID = "sandbox-20211130-kgd99i"
	topicID   = "raw-products-from-bucket"
)

func CreateNewClient() (*pubsub.Client, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		println("errors: ", err)
		return nil, err
	}
	return client, nil
}

func PublishRawProducts(client *pubsub.Client, messages product.Products) error {
	var wg sync.WaitGroup
	var totalErrors uint64

	ctx := context.Background()

	t := client.Topic(topicID)

	t.PublishSettings.NumGoroutines = 100 //test big value.

	for i := 0; i < len(messages); i++ {
		payload, err := json.Marshal(messages[i])
		if err != nil {
			return err
		}

		result := t.Publish(ctx, &pubsub.Message{Data: payload})

		wg.Add(1)
		go func(i int, res *pubsub.PublishResult) {
			defer wg.Done()
			_, err := res.Get(ctx)
			if err != nil {
				println("Failed to publish message")
				atomic.AddUint64(&totalErrors, 1)
				return
			}
		}(i, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, len(messages))
	}
	return nil
}
