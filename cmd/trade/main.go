package trade

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

	kafkaMessageChannel := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "mygroup",
		"auto.offset.reset": "earliest",
	}

	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewKafkaConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMessageChannel)

	book := entity.NewBook(ordersIn, ordersOut, waitGroup)

	go book.Trade()

	go func() {
		for message := range kafkaMessageChannel {
			waitGroup.Add(1)

			fmt.Println(message.Value)

			tradeInput := dto.TradeInput{}
			error := json.Unmarshal(message.Value, &tradeInput)

			if error != nil {
				panic(error)
			}

			order := transformer.TransformInput(tradeInput)

			ordersIn <- order
		}
	}()

	for result := range ordersOut {
		output := transformer.TransformOutput(result)

		outputJson, error := json.MarshalIndent(output, "", " ")

		if error != nil {
			panic(error)
		}

		producer.Publish(outputJson, []byte("orders"), "output")
	}
}
