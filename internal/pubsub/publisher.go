package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/AndriiShevchun/Parser/internal/product"
)

func PublishRawProduct(client *pubsub.Client, product []product.Product) error {
	//projectID := "sandbox-20211130-kgd99i"
	topicID := "raw-products-from-bucket"

	// msg := "Hello World"
	ctx := context.Background()

	t := client.Topic(topicID)
	t.PublishSettings.NumGoroutines = 1

	payload, err := json.Marshal(product)
	if err != nil {
		return err
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: payload,
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	println("Published a message; msg ID: %v\n", id)

	return nil
}

func PublishRawProducts(client *pubsub.Client, messages []product.Product) error {
	println(messages)
	var wg sync.WaitGroup
	var totalErrors uint64

	topicID := "raw-products-from-bucket"

	ctx := context.Background()

	t := client.Topic(topicID)

	for i := 0; i < len(messages); i++ {
		payload, err := json.Marshal(messages[i])
		if err != nil {
			return err
		}

		result := t.Publish(ctx, &pubsub.Message{Data: payload})

		wg.Add(1)
		go func(i int, res *pubsub.PublishResult) {
			defer wg.Done()
			id, err := res.Get(ctx)
			if err != nil {
				println("Failed to publish message")
				atomic.AddUint64(&totalErrors, 1)
				return
			}
			println("Published message %d; msg ID: %v\n", i, id)
		}(i, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, len(messages))
	}
	return nil
}
