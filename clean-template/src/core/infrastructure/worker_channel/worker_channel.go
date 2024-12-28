package worker_channel

import (
	"log"

	"github.com/nanda03dev/go-ms-template/src/common"
)

// Declare a global channel
var crudEventChannel = make(chan common.Event, 10000)

// Function to push data to the channel
func PushToCRUDChannel(event common.Event) {
	select {
	case crudEventChannel <- event:
		// Successfully pushed
	default:
		// Channel is full, log or handle overflow
		log.Println("Channel is full, dropping data for event:", event)
	}
}

// Function to get the channel
func GetCRUDEventChannel() chan common.Event {
	return crudEventChannel
}
