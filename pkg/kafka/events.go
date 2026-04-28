package kafka

import "time"

// Типы событий - константы, чтобы не писать строки руками
const (
	EventTaskCreated   = "TASK_CREATED"
	EventTaskUpdated   = "TASK_UPDATED"
	EventTaskDeleted   = "TASK_DELETED"
	EventTaskCompleted = "TASK_COMPLETED"

	EventListCreated = "LIST_CREATED"
	EventListUpdated = "LIST_UPDATED"
	EventListDeleted = "LIST_DELETED"

	EventUserRegistered = "USER_REGISTERED"
	EventUserLoggedIn   = "USER_LOGGED_IN"
)

// Event — универсальная структура любого события
type Event struct {
	EventID   string      `json:"event_id"`
	EventType string      `json:"event_type"`
	Timestamp time.Time   `json:"timestamp"`
	UserID    int         `json:"user_id"`
	ListID    int         `json:"list_id,omitempty"`
	TaskID    int         `json:"task_id,omitempty"`
	Data      interface{} `json:"data"`
}
