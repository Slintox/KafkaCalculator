package app

import (
	"context"
	"strconv"
	"strings"
	"time"

	"ExecutorService/config"
	"ExecutorService/pkg/kafka_client"
	"ExecutorService/pkg/kafka_client/consumer"
	"ExecutorService/pkg/kafka_client/producer"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logrus.Infof("%+v\n", *cfg)

	kafkaClient := kafka_client.NewKafka(&kafka_client.Config{
		URL: cfg.Kafka.URL,
	})

	expressionsConsumer := kafkaClient.NewConsumer(&consumer.Config{
		Brokers:   strings.Split(cfg.Kafka.Brokers, ","),
		Topic:     cfg.Kafka.Expressions.Topic,
		GroupID:   cfg.Kafka.Expressions.GroupID,
		Partition: cfg.Kafka.Expressions.Partition,
	})

	calculationsProducer := kafkaClient.NewProducer(&producer.Config{
		Topic: cfg.Kafka.Calculations.Topic,
	})

	logrus.Infoln(expressionsConsumer.Config())

	var err error

	var expressionMessage kafka.Message
	expressionMessage, err = expressionsConsumer.ReadMessage(context.Background())
	if err != nil {
		logrus.Errorf("app - Run - expressionConsumer.ReadMessage: %s\n", err)
		return
	}

	logrus.Infof("Input msg: %+v\n", expressionMessage)

	var expression int
	expression, err = strconv.Atoi(string(expressionMessage.Value))
	if err != nil {
		logrus.Errorf("app - Run - strconv.Atoi(expressionMessage.Value): %s", err)
	}
	logrus.Infoln("Input expression is:", expression)

	if expression == 42 {
		logrus.Infoln("Success! The question is 42!")
	}

	calculation := expression * 2

	calculationMessage := kafka.Message{
		Key:   expressionMessage.Key,
		Value: []byte(strconv.Itoa(calculation)),
		Time:  time.Now(),
	}

	err = calculationsProducer.WriteMessages(context.Background(), calculationMessage)
	if err != nil {
		logrus.Errorf("app - Run - expressionsProducer.WriteMessages: %s\n", err)
	}

	logrus.Infoln("Output calculation is:", calculation)
	logrus.Infof("Output msg: %+v\n", calculationMessage)

	err = expressionsConsumer.Close()
	if err != nil {
		logrus.Errorf("app - Run - producer.Close: %s", err)
	}

	err = calculationsProducer.Close()
	if err != nil {
		logrus.Errorf("app - Run - consumer.Close: %s", err)
	}

	logrus.Infoln("Exit 0")

	return
}
