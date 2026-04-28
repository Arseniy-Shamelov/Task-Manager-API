package events

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"todo-app/pkg/kafka"
)

// Handler — интерфейс для всех обработчиков событий
type Handler interface {
	Handle(event kafka.Event) error
}

// EventProcessor — управляет всеми consumer'ами
type EventProcessor struct {
	cfg         kafka.Config
	taskHandler Handler
	listHandler Handler
	userHandler Handler
}

func NewEventProcessor(cfg kafka.Config, task, list, user Handler) *EventProcessor {
	return &EventProcessor{
		cfg:         cfg,
		taskHandler: task,
		listHandler: list,
		userHandler: user,
	}
}

// Start — запускается в отдельной goroutine, читает все топики
func (p *EventProcessor) Start(ctx context.Context) {
	go p.consume(ctx, p.cfg.Topics.TaskEvents, p.taskHandler)
	go p.consume(ctx, p.cfg.Topics.ListEvents, p.listHandler)
	go p.consume(ctx, p.cfg.Topics.UserEvents, p.userHandler)
}

// consume — бесконечный цикл чтения одного топика
func (p *EventProcessor) consume(ctx context.Context, topic string, handler Handler) {
	consumer := kafka.NewConsumer(p.cfg, topic)
	defer consumer.Close()

	logrus.Infof("[EventProcessor] Started consuming topic: %s", topic)

	for {
		select {
		case <-ctx.Done():
			logrus.Infof("[EventProcessor] Stopped consuming topic: %s", topic)
			return
		default:
			data, err := consumer.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				logrus.Errorf("[EventProcessor] Read error on topic %s: %v", topic, err)
				continue
			}

			var event kafka.Event
			if err := json.Unmarshal(data, &event); err != nil {
				logrus.Errorf("[EventProcessor] Failed to unmarshal event: %v", err)
				continue
			}

			if err := handler.Handle(event); err != nil {
				logrus.Errorf("[EventProcessor] Handler error: event_type=%s, error=%v",
					event.EventType, err)
			}
		}
	}
}
