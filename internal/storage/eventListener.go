package storage

import (
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
	"strings"
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
	var data ObjectStorageData
	if err := e.DataAs(&data); err != nil {
		println("event.DataAs: %v", err)
		return ""
	}

	if strings.HasPrefix(data.Name, "data/") {
		log.Printf("File: %s recieved", data.Name)
		return data.Name
	} else {
		return ""
	}
}
