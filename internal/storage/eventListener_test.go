package storage_test

import (
	"errors"
	"github.com/AndriiShevchun/parse-pc-files/internal/storage"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListenStorageEvent(t *testing.T) {
	//ctx := context.Background()
	tests := []struct {
		name             string
		event            event.Event
		expectedFileName string
		expectedError    error
	}{
		{
			name: "File in 'data' folder",
			event: event.Event{
				Context:     &event.EventContextV1{ID: "UNIT TEST"},
				DataEncoded: []byte(`{"bucket": "test","name": "data/test_file.json"}`),
			},
			expectedFileName: "data/test_file.json",
			expectedError:    nil,
		},
		{
			name: "Any file in folder 'data' will dispatch",
			event: event.Event{
				Context:     &event.EventContextV1{ID: "UNIT TEST"},
				DataEncoded: []byte(`{"bucket": "test","name": "data/bla_bla.txt"}`),
			},
			expectedFileName: "data/bla_bla.txt",
			expectedError:    nil,
		},
		{
			name: "File without folder",
			event: event.Event{
				Context:     &event.EventContextV1{ID: "UNIT TEST"},
				DataEncoded: []byte(`{"bucket": "test","name": "test_file.json"}`),
			},
			expectedFileName: "",
			expectedError:    errors.New("file: test_file.json ignored"),
		},
		{
			name: "File in another folder",
			event: event.Event{
				Context:     &event.EventContextV1{ID: "UNIT TEST"},
				DataEncoded: []byte(`{"bucket": "test","name": "processed/test_file.json"}`),
			},
			expectedFileName: "",
			expectedError:    errors.New("file: processed/test_file.json ignored"),
		},
		{
			name: "Invalid event",
			event: event.Event{
				Context:     &event.EventContextV1{ID: "UNIT TEST"},
				DataEncoded: []byte(`{"bucket": 1,"name": "processed/test_file.json"}`),
			},
			expectedFileName: "",
			expectedError:    errors.New(`[json] found bytes "{"bucket": 1,"name": "processed/test_file.json"}", but failed to unmarshal: json: cannot unmarshal number into Go struct field ObjectStorageData.bucket of type string`),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := storage.ListenStorageEvent(test.event)
			assert.Equal(t, test.expectedFileName, f)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
