package interfaces

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	SentinelMetrics EventType = "sentinel:metrics"
	ProcessSignal   EventType = "process:signal" // TO BE IMPLEMENTED
	ProcessWatch    EventType = "process:watch"  // TO BE IMPLEMENTED
	VmException     EventType = "vm:exception"   // TO BE IMPLEMENTED
	SentinelStatus  EventType = "sentinel:status"
)

// EventPayload represents the payload of an event sent to the webhook
// it contains the timestamp of the event and the type of the event
// it also contains the data of the event that can be of different types
// if it's a vm:metrics event it contains the environment metrics of the vm and the processs metrics of the monitored processes
// if it's a process:signal event it contains the signal that was sent by a specific process to the sentinel through the Sentinel API
// if it's a process:watch event it contains information about the process that is watched by the sentinel for example if the process died or if the process is consuming too much memory
// if it's a vm:exception event it contains critical information about the vm that the sentinel is monitoring (where sentinel is running)
// if it's a sentinel:status event it contains information about the status of the sentinel (if it's running or not, when it started or stopped)
// the data field is an interface{} because the data can be of different types
type event struct {
	ID         string      `json:"id"`
	Timestamp  string      `json:"timestamp"`
	EventType  EventType   `json:"eventType"`
	SentinelID string      `json:"sentinelID"`
	Payload    interface{} `json:"payload"`
}

func NewEvent(sentinelID string, eventType EventType, data interface{}) *event {
	return &event{
		ID:         uuid.New().String(),
		Timestamp:  time.Now().Format(time.RFC3339),
		EventType:  eventType,
		SentinelID: sentinelID,
		Payload:    data,
	}
}
