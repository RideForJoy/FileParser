package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"strings"
)

const (
	bucket = "pc-file"
)

func CreateNewClient(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Errors happened during storage client creating: ", err)
	}
	return client
}

func ReadFile(ctx context.Context, client *storage.Client, path string) *storage.Reader {
	fmt.Println("Opening file: ", path)
	file, _ := client.Bucket(bucket).Object(path).NewReader(ctx)
	return file
}

//folderPath string

func MoveFile(ctx context.Context, client *storage.Client, folder string, movePath string, path string) error {
	dsdName := strings.Replace(path, "data", folder, 1)

	var fullPath string
	if movePath != "" {
		fullPath = strings.Replace(dsdName, "/", "/"+movePath+"/", 1)
	} else {
		fullPath = dsdName
	}

	src := client.Bucket(bucket).Object(path)
	dst := client.Bucket(bucket).Object(fullPath)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return fmt.Errorf("Object(%q).CopierFrom(%q).Run: %v", fullPath, path, err)
	}
	if err := src.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", path, err)
	}

	fmt.Printf("Blob %v moved to %v.\n", path, fullPath)

	return nil
}
