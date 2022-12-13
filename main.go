package main

import (
	"context"
	"github.com/AndriiShevchun/Parser/internal/product"
	"github.com/AndriiShevchun/Parser/internal/pubsub"
	"github.com/AndriiShevchun/Parser/internal/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
)

func init() {
	functions.CloudEvent("ProcessPCfile", ProcessPCfile)
}

func streamFile(path string) {
	pubSubClient, _ := pubsub.CreateNewClient()
	var i = 0
	var products []product.Product

	stream := product.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Status == "" {
				i++
				products = append(products, data.Product)
			}
			if i == 1000 || data.Status == "finish" {
				err := pubsub.PublishRawProducts(pubSubClient, products)
				if err != nil {
					return
				}
				println("we", i)
				i = 0
				products = nil
			}

			if data.Error != nil {
				log.Println(data.Error)
			}
		}
	}()
	stream.Start(path)

}

/*
func main() {
	println(time.Now().String())
	//fileName := "Takeoff_product_catalog_full_D02_20221207230258.json"

	//if storage.CopyFromBucketToLocal(fileName) {
	streamFile("file.json")
	//}
}
*/

func ProcessPCfile(ctx context.Context, e event.Event) error {

	fileName := storage.ListenStorageEvent(e)

	if storage.CopyFromBucketToLocal(fileName) {
		streamFile("file.json")
	}
	return nil
}
