package storage

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
)

type ObjectStorageData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration int64     `json:"metageneration,string,omitempty"`
	TimeCreated    time.Time `json:"timeCreated,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
}

func ListenStorageEvent(e event.Event) (string, error) {
	var data ObjectStorageData
	if err := e.DataAs(&data); err != nil {
		return "", err
	}

	if strings.HasPrefix(data.Name, "data/") {
		log.Printf("File: %s recieved", data.Name)
		return data.Name, nil
	} else {
		errorString := "file: " + data.Name + " ignored"
		return "", errors.New(errorString)
	}
}
