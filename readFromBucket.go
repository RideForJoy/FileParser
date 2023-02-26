package read

import (
	gcpStorage "cloud.google.com/go/storage"
	"context"
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
	"time"

	"github.com/AndriiShevchun/parse-pc-files/internal/product"
	"github.com/AndriiShevchun/parse-pc-files/internal/pubsub"
	"github.com/AndriiShevchun/parse-pc-files/internal/storage"
)

func processFile(ctx context.Context, storageClient *gcpStorage.Client, fileName string) error {
	pubSubClient, _ := pubsub.CreateNewClient()
	defer pubSubClient.Close()

	jsonFile := storage.Read(ctx, storageClient, fileName)
	products, err := product.Unmarshal(jsonFile)
	if err != nil {
		return err
	}
	chunkedProducts := product.ChunkBy(products, 5000) //value can be smaller/bigger based on function memory

	ProcessingStartTime := time.Now()
	println("Raw_products publishing started at: ", ProcessingStartTime.String())

	for i := 0; i < len(chunkedProducts); i++ {
		//publishStart := time.Now()
		err := pubsub.PublishRawProducts(pubSubClient, chunkedProducts[i])
		if err != nil {
			return err
		}
		//now := time.Now()
		//fmt.Println("Published chunk №", i, "with length ", len(chunkedProducts[i]), "in ", now.Sub(publishStart).String())
	}
	finish := time.Now()

	println("Raw_products publishing finished at: ", finish.Sub(ProcessingStartTime).String(), "for count: ", len(products))
	return nil
}

func ListenForFile(ctx context.Context, e event.Event) error {
	storageClient, err := storage.NewStorageClient(ctx)
	if err != nil {
		println("Can't start storageClient due to error: ", err)
		return nil
	}
	defer storageClient.Close()

	fileName, err := storage.ListenStorageEvent(e)
	if err != nil {
		println(err)
		return nil
	}
	if storage.IsPCv6(fileName) {
		log.Printf("File: %s dispathed as PCv6 structure", fileName)
		fileDate, fileDateErr := storage.DateFromName(fileName)
		if fileDateErr != nil {
			if err := storage.Move(ctx, storageClient, "unknown", "", fileName); err != nil {
				println("Can't move file. Function stop due to error: ", err)
			}
			return nil
		}

		movePath := storage.MovePath(fileDate)
		if storage.IsFresh(fileDate) {
			if err := processFile(ctx, storageClient, fileName); err != nil {
				println("Can't process file due to error: ", err)
			}

			if err := storage.Move(ctx, storageClient, "processed", movePath, fileName); err != nil {
				println("Can't move file due to error: ", err)
			}

			println("File: " + fileName + " processed and moved to 'processed' folder")
		} else {
			if err := storage.Move(ctx, storageClient, "skipped", movePath, fileName); err != nil {
				println("Can't move file due to error: ", err)
			}
			println("File: " + fileName + " moved to 'skipped' folder")
		}
	} else {
		println("File: " + fileName + " is ignored. Should be processed by ETL-A")
	}

	return nil
}
