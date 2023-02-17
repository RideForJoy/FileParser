package read

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"time"

	"github.com/AndriiShevchun/parse-pc-files/internal/product"
	"github.com/AndriiShevchun/parse-pc-files/internal/pubsub"
	"github.com/AndriiShevchun/parse-pc-files/internal/storage"
)

func ProcessFile(ctx context.Context, e event.Event) error {
	pubSubClient, _ := pubsub.CreateNewClient()

	fileName := storage.ListenStorageEvent(e)
	jsonFile := storage.ReadFile(fileName)
	products := product.Unmarshal(jsonFile)
	chunkedProducts := product.ChunkBy(products, 5000) //value can be smaller/bigger based on function memory

	ProcessingStartTime := time.Now()
	println("File processing started at: ", ProcessingStartTime.String())

	for i := 0; i < len(chunkedProducts); i++ {
		publishStart := time.Now()
		err := pubsub.PublishRawProducts(pubSubClient, chunkedProducts[i])
		if err != nil {
			return nil
		}
		now := time.Now()
		fmt.Println("Published chunk №", i, "with length ", len(chunkedProducts[i]), "in ", now.Sub(publishStart).String())
	}
	finish := time.Now()

	println("File processing finished at: ", finish.Sub(ProcessingStartTime).String())
	return nil

}
