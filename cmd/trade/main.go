package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/eng-gabrielscardoso/exchange-core/app/infra/kafka"
	"github.com/eng-gabrielscardoso/exchange-core/app/market/dto"
	"github.com/eng-gabrielscardoso/exchange-core/app/market/entity"
	"github.com/eng-gabrielscardoso/exchange-core/app/market/transformer"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	waitGroup := &sync.WaitGroup{}

	defer waitGroup.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	}

	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewKafkaConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan)

	book := entity.NewBook(ordersIn, ordersOut, waitGroup)

	go book.Trade()

	go func() {
		for msg := range kafkaMsgChan {
			waitGroup.Add(1)

			fmt.Println(string(msg.Value))

			tradeInput := dto.TradeInput{}

			err := json.Unmarshal(msg.Value, &tradeInput)

			if err != nil {
				panic(err)
			}

			order := transformer.TransformInput(tradeInput)

			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)

		outputJson, err := json.MarshalIndent(output, "", "  ")

		fmt.Println(string(outputJson))

		if err != nil {
			fmt.Println(err)
		}

		producer.Publish(outputJson, []byte("orders"), "output")
	}
}
