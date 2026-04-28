package events

import (
	"github.com/sirupsen/logrus"
	"todo-app/pkg/kafka"
)

type UserEventHandler struct{}

func NewUserEventHandler() *UserEventHandler {
	return &UserEventHandler{}
}

func (h *UserEventHandler) Handle(event kafka.Event) error {
	logrus.Infof("[UserHandler] Processing event: type=%s, event_id=%s",
		event.EventType, event.EventID)

	switch event.EventType {
	case kafka.EventUserRegistered:
		return h.handleUserRegistered(event)
	case kafka.EventUserLoggedIn:
		return h.handleUserLoggedIn(event)
	default:
		logrus.Warnf("[UserHandler] Unknown event type: %s", event.EventType)
		return nil
	}
}

func (h *UserEventHandler) handleUserRegistered(event kafka.Event) error {
	logrus.Infof("[UserHandler] User registered: user_id=%d", event.UserID)
	// TODO: отправить приветственный email
	return nil
}

func (h *UserEventHandler) handleUserLoggedIn(event kafka.Event) error {
	logrus.Infof("[UserHandler] User logged in: user_id=%d", event.UserID)
	// TODO: обновить last_seen, проверить подозрительную активность
	return nil
}

var _ Handler = (*UserEventHandler)(nil)
