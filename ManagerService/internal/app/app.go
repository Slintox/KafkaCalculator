package app

import (
	"context"
	"strconv"
	"strings"
	"time"

	"ManagerService/config"
	"ManagerService/pkg/kafka_client"
	"ManagerService/pkg/kafka_client/consumer"
	"ManagerService/pkg/kafka_client/producer"

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

	expressionsProducer := kafkaClient.NewProducer(&producer.Config{
		Topic: cfg.Kafka.Expressions.Topic,
	})

	calculationsConsumer := kafkaClient.NewConsumer(&consumer.Config{
		Brokers:   strings.Split(cfg.Kafka.Brokers, ","),
		Topic:     cfg.Kafka.Calculations.Topic,
		GroupID:   cfg.Kafka.Calculations.GroupID,
		Partition: cfg.Kafka.Calculations.Partition,
	})

	Expression := "42"

	expressionsCount := 0
	expressionMessage := kafka.Message{
		Key:   []byte(strconv.Itoa(expressionsCount)),
		Value: []byte(Expression),
		Time:  time.Now(),
	}

	expectedCalculation := 42

	logrus.Infoln("Output expression is:", Expression)
	logrus.Infof("Output msg: %+v\n", expressionMessage)

	err := expressionsProducer.WriteMessages(context.Background(), expressionMessage)
	if err != nil {
		logrus.Errorf("app - Run - expressionsProducer.WriteMessages: %s", err)
		return
	}

	time.Sleep(time.Second * 2)

	var calculationMessage kafka.Message
	calculationMessage, err = calculationsConsumer.ReadMessage(context.Background())
	if err != nil {

		logrus.Errorf("app - Run - calculationConsumer.ReadMessage: %s\n", err)
		return
	}

	var calculation int
	calculation, err = strconv.Atoi(string(calculationMessage.Value))
	if err != nil {
		logrus.Errorf("app - Run - strconv.Atoi(calculationMessage.Value): %s", err)
	}

	logrus.Infoln("Input calculation is:", calculation)
	logrus.Infof("Input msg: %+v\n", calculationMessage)

	if calculation == expectedCalculation {
		logrus.Infof("Success! The answer is 42! (%d)\n", calculation)
		//logrus.Infoln("Success! The answer is 42!")
	} else {
		logrus.Infof("%+v\n", calculationMessage)
	}

	err = expressionsProducer.Close()
	if err != nil {
		logrus.Errorf("app - Run - producer.Close: %s\n", err)
	}

	err = calculationsConsumer.Close()
	if err != nil {
		logrus.Errorf("app - Run - consumer.Close: %s\n", err)
	}

	logrus.Infoln("Exit 0")

	return
}
