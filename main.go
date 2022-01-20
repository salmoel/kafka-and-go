package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	TopicName = "batata"
	Address   = "localhost:9092"
	Key = 1234556
)

func main() {
	cm := kafka.ConfigMap{
		"bootstrap.servers": Address,
		"go.logs.channel.enable": true,
	}

	p, err := kafka.NewProducer(&cm)
	if err != nil {
		panic(err.Error())
	}

	adm, err := kafka.NewAdminClientFromProducer(p)
	if err != nil {
		panic(err.Error())
	}

	tps, err := adm.CreateTopics(context.TODO(), []kafka.TopicSpecification{
		{
			Topic: TopicName,
			NumPartitions: 1,
			ReplicationFactor: 1,
		},
	})
	if err != nil {
		panic(err.Error())
	}

	for _, tp := range tps {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", tp)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &tp.Topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte("eeeeeeita"),
			Key: []byte("oooooxe"),
		}, nil)
	}
}
