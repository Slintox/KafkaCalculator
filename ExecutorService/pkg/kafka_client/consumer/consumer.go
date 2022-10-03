package consumer

import "github.com/segmentio/kafka-go"

type Config struct {
	GroupID   string
	Topic     string
	Partition int
	Brokers   []string
}

type Consumer struct {
	*kafka.Reader
}
