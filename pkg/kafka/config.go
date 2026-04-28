package kafka

import "time"

type Config struct {
	Brokers  []string
	GroupID  string
	Topics   Topics
	Producer ProducerConfig
	Consumer ConsumerConfig
}

type Topics struct {
	TaskEvents string
	ListEvents string
	UserEvents string
}

type ProducerConfig struct {
	RetryMax     int
	BatchSize    int
	WriteTimeout time.Duration
}

type ConsumerConfig struct {
	SessionTimeout time.Duration
	MaxPollRecords int
}
