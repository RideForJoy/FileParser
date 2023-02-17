package storage

import (
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
	"time"
)

type ObjectStorageData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration int64     `json:"metageneration,string,omitempty"`
	TimeCreated    time.Time `json:"timeCreated,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
}

func ListenStorageEvent(e event.Event) string {

	log.Printf("Event ID: %s", e.ID())
	log.Printf("Event Type: %s", e.Type())

	var data ObjectStorageData
	if err := e.DataAs(&data); err != nil {
		println("event.DataAs: %v", err)
		return ""
	}

	log.Printf("Bucket: %s", data.Bucket)
	log.Printf("File: %s", data.Name)
	log.Printf("Metageneration: %d", data.Metageneration)
	log.Printf("Created: %s", data.TimeCreated)
	log.Printf("Updated: %s", data.Updated)

	return data.Name
}
