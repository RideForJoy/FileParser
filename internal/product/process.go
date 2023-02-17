package product

import (
	"encoding/json"
	"fmt"
	"io"
)

func Unmarshal(file io.Reader) Products {
	byteValue, _ := io.ReadAll(file)

	var products Products

	err := json.Unmarshal(byteValue, &products)
	if err != nil {
		fmt.Println("Error during unmarshal products: ", err) //TODO add better error handling
	}
	return products
}

func ChunkBy(products []Product, chunkSize int) (chunks [][]Product) {
	for chunkSize < len(products) {
		products, chunks = products[chunkSize:], append(chunks, products[0:chunkSize:chunkSize])
	}
	return append(chunks, products)
}
