package product_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/RideForJoy/parse-pc-files/internal/product"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name            string
		file            string
		expectedProduct product.Product
		expectedError   error
	}{
		{
			name: "full product",
			file: `[{"product-id":"banana", "barcodes":["0005680025047"], "corp-ids":[], "ecom-ids":["1"], "description":"Banana product", "name":"Fresh banana", "temperature-zone":["ambient"], "image":"www.nga-herzog.org"}]`, //,"retail-item":{"weight":{"unit-of-measure":"KG","weight":2.566},"dimensions":{"unit-of-measure":"CM","height":11.684,"width":38.1,"length":12.827}}}]`,
			expectedProduct: product.Product{
				TomID:           "banana",
				Barcodes:        []string{"0005680025047"},
				CorpIds:         []string{},
				EcomIds:         []string{"1"},
				Description:     "Banana product",
				Name:            "Fresh banana",
				TemperatureZone: []string{"ambient"},
				Image:           "www.nga-herzog.org",
				//RetailItem: {Dimensions: {UnitOfMeasure: "CM", Height: 11.684, Width: 38.1, Length: 12.827},
				//	Weight: {UnitOfMeasure: "KG", Weight: 2.566}},

				UnknownFields: map[string]json.RawMessage{}},
			expectedError: nil,
		},

		{
			name: "With 1 field",
			file: `[{"product-id":"p1"}]`,
			expectedProduct: product.Product{
				TomID:         "p1",
				UnknownFields: map[string]json.RawMessage{}},
			expectedError: nil,
		},
		{
			name: "With additional fields",
			file: `[{"product-id":"p1", "additional-fields": "Additional fields value"}]`,
			expectedProduct: product.Product{
				TomID:         "p1",
				UnknownFields: map[string]json.RawMessage{"additional-fields": json.RawMessage(`"Additional fields value"`)}},
			expectedError: nil,
		},
		{
			name:            "invalid file",
			file:            `[{"product-id":"p1"`,
			expectedProduct: product.Product{},
			expectedError:   errors.New("errors during unmarshal products: unexpected end of JSON input"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buffer bytes.Buffer
			buffer.WriteString(test.file)
			r, err := product.Unmarshal(&buffer)

			assert.Equal(t, test.expectedError, err)
			if err == nil {
				assert.Equal(t, test.expectedProduct, r[0])
			}
		})
	}
}

func TestChunkBy(t *testing.T) {
	p3 := product.Products{
		{
			TomID: "1",
		},
		{
			TomID: "2",
		},
		{
			TomID: "3",
		},
	}

	p1 := product.Products{
		{
			TomID: "1",
		},
	}

	tests := []struct {
		name                 string
		products             []product.Product
		chunkSize            int
		expectedChunksLength int
	}{
		{
			name:                 "3 products split into 2 chunk with 2 and 1 product",
			products:             p3,
			chunkSize:            2,
			expectedChunksLength: 2,
		},
		{
			name:                 "1 product can be split only to 1 chunk regardless of chunkSize",
			products:             p1,
			chunkSize:            2,
			expectedChunksLength: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedChunksLength, len(product.ChunkBy(test.products, test.chunkSize)))
		})
	}
}
