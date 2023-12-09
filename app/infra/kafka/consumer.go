package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// TYPE DECLARATIONS

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

// CONSTRUCTORS

func NewKafkaConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

// METHODS

func (c *Consumer) Consume(messageChannel chan *ckafka.Message) error {
	consumer, error := ckafka.NewConsumer(c.ConfigMap)

	if error != nil {
		panic(error)
	}

	error = consumer.SubscribeTopics(c.Topics, nil)

	if error != nil {
		panic(error)
	}

	for {
		message, error := consumer.ReadMessage(-1)

		if error == nil {
			messageChannel <- message
		}
	}
}
