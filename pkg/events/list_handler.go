package events

import (
	"github.com/sirupsen/logrus"
	"todo-app/pkg/kafka"
)

type ListEventHandler struct{}

func NewListEventHandler() *ListEventHandler {
	return &ListEventHandler{}
}

func (h *ListEventHandler) Handle(event kafka.Event) error {
	logrus.Infof("[ListHandler] Processing event: type=%s, event_id=%s",
		event.EventType, event.EventID)

	switch event.EventType {
	case kafka.EventListCreated:
		return h.handleListCreated(event)
	case kafka.EventListUpdated:
		return h.handleListUpdated(event)
	case kafka.EventListDeleted:
		return h.handleListDeleted(event)
	default:
		logrus.Warnf("[ListHandler] Unknown event type: %s", event.EventType)
		return nil
	}
}

func (h *ListEventHandler) handleListCreated(event kafka.Event) error {
	logrus.Infof("[ListHandler] List created: list_id=%d, user_id=%d",
		event.ListID, event.UserID)
	return nil
}

func (h *ListEventHandler) handleListUpdated(event kafka.Event) error {
	logrus.Infof("[ListHandler] List updated: list_id=%d, user_id=%d",
		event.ListID, event.UserID)
	return nil
}

func (h *ListEventHandler) handleListDeleted(event kafka.Event) error {
	logrus.Infof("[ListHandler] List deleted: list_id=%d, user_id=%d",
		event.ListID, event.UserID)
	return nil
}

var _ Handler = (*ListEventHandler)(nil)
