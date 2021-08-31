package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()

	Publish(
		fmt.Sprintf("[%v]: Message", time.Now().String()),
		"goteste",
		producer,
		nil,
		deliveryChan,
	)

	Publish(
		fmt.Sprintf("[%v]: Message with key", time.Now().String()),
		"goteste",
		producer,
		[]byte("key"), // This message will always go to the same partition
		deliveryChan,
	)

	go DeliveryReport(deliveryChan)

	// e := <-deliveryChan
	// msg := e.(*kafka.Message)

	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Error on sending")
	// } else {
	// 	fmt.Println("Message sent: ", msg.TopicPartition)
	// }

	producer.Flush(10000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "full-cycle-20-kafka_kafka_1:9092",
		"delivery.timeout.ms": "0",     // Wait until delivered
		"acks":                "all",   // "0" no guarantee, "1" leader guarantee and "all" everyone
		"enable.idempotence":  "false", // If true, "acks" need to be "all"
	}
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return producer
}

func Publish(
	msg string,
	topic string,
	producer *kafka.Producer,
	key []byte,
	deliveryChan chan kafka.Event,
) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Error on sending")
			} else {
				fmt.Println("Message sent: ", ev.TopicPartition)
			}
		}
	}
}
