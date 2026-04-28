package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// Интерфейс
type Consumer interface {
	ReadMessage(ctx context.Context) ([]byte, error)
	Close() error
}

// Реализация
type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewConsumer(cfg Config, topic string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		GroupID:  cfg.GroupID,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6, // 10MB
	})

	logrus.Infof("[Kafka Consumer] Created consumer for topic: %s, group: %s", topic, cfg.GroupID)

	return &KafkaConsumer{reader: reader}
}

func (c *KafkaConsumer) ReadMessage(ctx context.Context) ([]byte, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	logrus.Infof("[Kafka Consumer] Message consumed: topic=%s, offset=%d",
		msg.Topic, msg.Offset)

	return msg.Value, nil
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
