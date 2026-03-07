package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

type Metrics struct {
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
	Vibration   float64 `json:"vibration"`
}

type Metadata struct {
	Location        string `json:"location"`
	FirmwareVersion string `json:"firmware_version"`
}

type Event struct {
	EventID   string `json:"event_id"`
	DeviceID  string `json:"device_id"`
	Timestamp int64  `json:"timestamp"`

	Metrics  Metrics   `json:"metrics"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

func generateMetrics() Metrics {
	return Metrics{
		Temperature: 20 + rand.Float64()*10,
		Pressure:    1 + rand.Float64()*0.2,
		Vibration:   rand.Float64(),
	}
}

func generateMetadata() *Metadata {
	return &Metadata{
		Location:        fmt.Sprintf("factory-floor-%d", rand.IntN(10)),
		FirmwareVersion: fmt.Sprintf("%d.%d.%d", rand.IntN(4), rand.IntN(10), rand.IntN(10)),
	}
}

func simulate_device(deviceID string, out chan<- Event) {
	for {
		var metadata *Metadata
		if rand.IntN(2) == 0 {
			metadata = generateMetadata()
		}
		event := Event{
			EventID:   uuid.NewString(),
			DeviceID:  deviceID,
			Timestamp: time.Now().UnixMilli(),
			Metrics:   generateMetrics(),
			Metadata:  metadata,
		}

		out <- event
		time.Sleep(time.Second)
	}

}

func main() {
	eventChan := make(chan Event, 1000)

	for i := range 100 {
		deviceID := fmt.Sprintf("device-%d", i)
		go simulate_device(deviceID, eventChan)
	}

	for event := range eventChan {
		fmt.Printf("%+v\n", event)
	}
}
