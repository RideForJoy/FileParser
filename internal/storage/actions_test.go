package storage_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/AndriiShevchun/parse-pc-files/internal/storage"
)

func TestNewStorageClient(t *testing.T) {
	ctx := context.Background()
	_, err := storage.NewStorageClient(ctx)
	assert.NoError(t, err)
}
