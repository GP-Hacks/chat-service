package servie_provider

import (
	"github.com/GP-Hacks/chat/internal/config"
	"github.com/GP-Hacks/chat/internal/infrastructure/broker"
)

func (s *ServiceProvider) KafkaConsumer() *broker.KafkaConsumer {
	if s.kafkaConsumer == nil {
		s.kafkaConsumer = broker.NewKafkaConsumer(config.Cfg.Kafka.Broker, config.Cfg.Kafka.Topic, config.Cfg.Kafka.GroupID, s.ChatService())
	}

	return s.kafkaConsumer
}

func (s *ServiceProvider) KafkaProducer() *broker.KafkaProducer {
	if s.kafkaProducer == nil {
		s.kafkaProducer = broker.NewKafkaProducer(config.Cfg.Kafka.Broker, config.Cfg.Kafka.Topic)
	}

	return s.kafkaProducer
}
