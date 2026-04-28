package events

import (
	"github.com/sirupsen/logrus"
	"todo-app/pkg/kafka"
)

type TaskEventHandler struct{}

func NewTaskEventHandler() *TaskEventHandler {
	return &TaskEventHandler{}
}

// Handle — точка входа, маршрутизирует по типу события
func (h *TaskEventHandler) Handle(event kafka.Event) error {
	logrus.Infof("[TaskHandler] Processing event: type=%s, event_id=%s",
		event.EventType, event.EventID)

	switch event.EventType {
	case kafka.EventTaskCreated:
		return h.handleTaskCreated(event)
	case kafka.EventTaskUpdated:
		return h.handleTaskUpdated(event)
	case kafka.EventTaskDeleted:
		return h.handleTaskDeleted(event)
	case kafka.EventTaskCompleted:
		return h.handleTaskCompleted(event)
	default:
		logrus.Warnf("[TaskHandler] Unknown event type: %s", event.EventType)
		return nil
	}
}

func (h *TaskEventHandler) handleTaskCreated(event kafka.Event) error {
	logrus.Infof("[TaskHandler] Task created: task_id=%d, user_id=%d, list_id=%d",
		event.TaskID, event.UserID, event.ListID)
	// TODO: отправить уведомление, обновить кэш и т.д.
	return nil
}

func (h *TaskEventHandler) handleTaskUpdated(event kafka.Event) error {
	logrus.Infof("[TaskHandler] Task updated: task_id=%d, user_id=%d",
		event.TaskID, event.UserID)
	return nil
}

func (h *TaskEventHandler) handleTaskDeleted(event kafka.Event) error {
	logrus.Infof("[TaskHandler] Task deleted: task_id=%d, user_id=%d",
		event.TaskID, event.UserID)
	return nil
}

func (h *TaskEventHandler) handleTaskCompleted(event kafka.Event) error {
	logrus.Infof("[TaskHandler] Task completed: task_id=%d, user_id=%d",
		event.TaskID, event.UserID)
	// TODO: отправить поздравление, обновить статистику
	return nil
}

// Проверка что TaskEventHandler реализует интерфейс Handler
var _ Handler = (*TaskEventHandler)(nil)
