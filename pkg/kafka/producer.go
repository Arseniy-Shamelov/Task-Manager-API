package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// Интерфейс — контракт, что умеет Producer
type Producer interface {
	Publish(topic string, event Event) error
	Close() error
}

// Реализация
type KafkaProducer struct {
	brokers []string
	timeout time.Duration
}

func NewProducer(cfg Config) *KafkaProducer {
	return &KafkaProducer{
		brokers: cfg.Brokers,
		timeout: 10 * time.Second,
	}
}

func (p *KafkaProducer) Publish(topic string, event Event) error {
	// Создаём writer (соединение) для конкретного topic
	w := &kafka.Writer{
		Addr:         kafka.TCP(p.brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: p.timeout,
		MaxAttempts:  3,
	}
	defer w.Close()

	// Сериализуем событие в JSON
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Ключ сообщения — user_id (гарантирует порядок событий одного юзера)
	key := []byte(fmt.Sprintf("%d", event.UserID))

	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: data,
	})
	if err != nil {
		logrus.Errorf("[Kafka Producer] Publish failed: topic=%s, error=%v", topic, err)
		return err
	}

	logrus.Infof("[Kafka Producer] Event published: topic=%s, event_type=%s, user_id=%d",
		topic, event.EventType, event.UserID)
	return nil
}

func (p *KafkaProducer) Close() error {
	return nil
}
