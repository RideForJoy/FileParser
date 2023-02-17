package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"os"
)

func CopyFromBucketToLocal(objectName string) bool {

	bucket := "product-catalog-file"
	//objectName := "Takeoff_product_catalog_full_D02_20221207230258.json" //"test-finalize.json"
	destFileName := "file1.json"

	f, err := os.Create(destFileName)
	if err != nil {
		println("os.Create: %v", err)
		return false
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("errors happened: ")
		// TODO: handle error.
	}

	fmt.Println("Opening file: ")
	// Open file to read.
	file, err := client.Bucket(bucket).Object(objectName).NewReader(ctx)
	defer file.Close()

	if _, err := io.Copy(f, file); err != nil {
		fmt.Println("errors happened: ")
		return false
	}

	fmt.Printf("Blob %v downloaded to local file %v\n\n", objectName, destFileName)
	return true
}
