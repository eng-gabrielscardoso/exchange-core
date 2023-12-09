package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// TYPE DECLARATIONS

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

// CONSTRUCTORS

func NewKafkaProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{
		ConfigMap: configMap,
	}
}

// METHODS

func (p *Producer) Publish(message interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)

	if err != nil {
		return err
	}

	kafkaMessage := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Key:            key,
		Value:          message.([]byte),
	}

	err = producer.Produce(kafkaMessage, nil)

	if err != nil {
		return err
	}

	return nil
}
