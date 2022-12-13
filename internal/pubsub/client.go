package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

func CreateNewClient() (*pubsub.Client, error) {
	projectID := "sandbox-20211130-kgd99i"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		println("errors: ", err)
		return nil, err
	}
	return client, nil
}
