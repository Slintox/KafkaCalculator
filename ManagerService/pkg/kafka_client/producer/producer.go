package producer

import (
	"github.com/segmentio/kafka-go"
)

type Config struct {
	Topic string
}

type Producer struct {
	*kafka.Writer
}
