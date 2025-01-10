package interfaces

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	SentinelMetrics EventType = "sentinel:metrics"
	ProcessSignal   EventType = "process:signal"   // TO BE IMPLEMENTED
	ProcessWatch    EventType = "process:watch"    // TO BE IMPLEMENTED
	SystemException EventType = "system:exception" // TO BE IMPLEMENTED
	SentinelStatus  EventType = "sentinel:status"  // TO BE IMPLEMENTED
)

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
