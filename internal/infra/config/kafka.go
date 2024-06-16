package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

func NewKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "balance",
	}
}
