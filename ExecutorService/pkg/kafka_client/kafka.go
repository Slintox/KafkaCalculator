package kafka_client

import (
	"ExecutorService/pkg/kafka_client/consumer"
	"ExecutorService/pkg/kafka_client/producer"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	URL     string
	Brokers string
}

type KafkaClient struct {
	Config     *Config
	Connection *kafka.Conn
}

func NewKafka(cfg *Config) *KafkaClient {
	k := &KafkaClient{
		Config: cfg,
	}

	return k
}

func (k *KafkaClient) NewProducer(cfg *producer.Config) *producer.Producer {
	prod := &producer.Producer{
		Writer: &kafka.Writer{
			Addr:                   kafka.TCP(k.Config.URL),
			Topic:                  cfg.Topic,
			Balancer:               nil,
			MaxAttempts:            0,
			BatchSize:              0,
			BatchBytes:             0,
			BatchTimeout:           0,
			ReadTimeout:            0,
			WriteTimeout:           0,
			RequiredAcks:           0,
			Async:                  false,
			Completion:             nil,
			Compression:            0,
			Logger:                 nil,
			ErrorLogger:            nil,
			Transport:              nil,
			AllowAutoTopicCreation: false,
		},
	}

	return prod
}

func (k *KafkaClient) NewConsumer(cfg *consumer.Config) *consumer.Consumer {
	readerConfig := kafka.ReaderConfig{
		Brokers:                cfg.Brokers,
		GroupID:                cfg.GroupID,
		GroupTopics:            nil,
		Topic:                  cfg.Topic,
		Partition:              cfg.Partition,
		Dialer:                 nil,
		QueueCapacity:          0,
		MinBytes:               10e3,
		MaxBytes:               10e6,
		MaxWait:                0,
		ReadLagInterval:        0,
		GroupBalancers:         nil,
		HeartbeatInterval:      0,
		CommitInterval:         0,
		PartitionWatchInterval: 0,
		WatchPartitionChanges:  false,
		SessionTimeout:         0,
		RebalanceTimeout:       0,
		JoinGroupBackoff:       0,
		RetentionTime:          0,
		StartOffset:            0,
		ReadBackoffMin:         0,
		ReadBackoffMax:         0,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         0,
		MaxAttempts:            0,
		OffsetOutOfRangeError:  false,
	}

	reader := kafka.NewReader(readerConfig)

	cons := &consumer.Consumer{
		Reader: reader,
	}

	return cons
}
