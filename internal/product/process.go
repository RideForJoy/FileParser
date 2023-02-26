package product

import (
	"encoding/json"
	"fmt"
	"io"
)

func Unmarshal(file io.Reader) (Products, error) {
	byteValue, _ := io.ReadAll(file)

	var products Products
	err := json.Unmarshal(byteValue, &products)
	if err != nil {
		return nil, fmt.Errorf("errors during unmarshal products: %v", err)
	}

	return products, nil
}

func ChunkBy(products []Product, chunkSize int) (chunks [][]Product) {
	for chunkSize < len(products) {
		products, chunks = products[chunkSize:], append(chunks, products[0:chunkSize:chunkSize])
	}
	return append(chunks, products)
}
