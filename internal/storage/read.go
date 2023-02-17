package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
)

const (
	bucket = "pc-file"
)

func ReadFile(path string) *storage.Reader {

	ctx := context.Background() //faced with "context canceled" error when pass
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("errors happened: ", err)
		// TODO: handle error.
	}

	fmt.Println("Opening file: ")
	// Open file to read.
	file, err := client.Bucket(bucket).Object(path).NewReader(ctx)

	return file
}
