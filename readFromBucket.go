package read

import (
	gcpStorage "cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
	"time"

	"github.com/AndriiShevchun/parse-pc-files/internal/file"
	"github.com/AndriiShevchun/parse-pc-files/internal/product"
	"github.com/AndriiShevchun/parse-pc-files/internal/pubsub"
	"github.com/AndriiShevchun/parse-pc-files/internal/storage"
)

func processFile(ctx context.Context, storageClient *gcpStorage.Client, fileName string) error {
	pubSubClient, _ := pubsub.CreateNewClient()
	defer pubSubClient.Close()

	jsonFile := storage.ReadFile(ctx, storageClient, fileName)
	products := product.Unmarshal(jsonFile)
	chunkedProducts := product.ChunkBy(products, 5000) //value can be smaller/bigger based on function memory

	ProcessingStartTime := time.Now()
	println("Raw_products publishing started at: ", ProcessingStartTime.String())

	for i := 0; i < len(chunkedProducts); i++ {
		publishStart := time.Now()
		err := pubsub.PublishRawProducts(pubSubClient, chunkedProducts[i])
		if err != nil {
			return err
		}
		now := time.Now()
		fmt.Println("Published chunk №", i, "with length ", len(chunkedProducts[i]), "in ", now.Sub(publishStart).String())
	}
	finish := time.Now()

	println("Raw_products publishing finished at: ", finish.Sub(ProcessingStartTime).String())
	return nil
}

func ListenForFile(ctx context.Context, e event.Event) error {
	storageClient := storage.CreateNewClient(ctx)
	defer storageClient.Close()

	fileName := storage.ListenStorageEvent(e)
	if fileName != "" && file.IsPCv6(fileName) {
		log.Printf("File: %s dispathed as PCv6 structure", fileName)
		fileDate, fileDateErr := file.DateFromName(fileName)
		if fileDateErr != nil {
			if err := storage.MoveFile(ctx, storageClient, "unknown", "", fileName); err != nil {
				println("Can't move file. Function stop due to error: ", err)
				return nil
			}
		}
		movePath := file.MovePath(fileDate)
		if file.IsFresh(fileDate) {
			err := processFile(ctx, storageClient, fileName)
			if err != nil {
				return nil
			}
			if err := storage.MoveFile(ctx, storageClient, "processed", movePath, fileName); err != nil {
				println("Can't move file. Function stop due to error: ", err)
				return nil
			}
			println("File: " + fileName + "processed and  moved to 'processed' folder")
		} else {
			if err := storage.MoveFile(ctx, storageClient, "skipped", movePath, fileName); err != nil {
				println("Can't move file. Function stop due to error: ", err)
				return nil
			}
			println("File: " + fileName + " moved to 'skipped' folder")
		}
	} else {
		println("File: " + fileName + " is ignored. Should be processed be ETL-A")
	}

	return nil

}
