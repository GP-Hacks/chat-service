package broker

import (
	"context"
	"encoding/json"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/GP-Hacks/chat/internal/services/chat_service"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader  *kafka.Reader
	service *chat_service.ChatService
}

type messageDto struct {
	AuthToken string `json:"auth_token,omitempty"`
	Content   string `json:"content,omitempty"`
}

func NewKafkaConsumer(broker, topic, groupID string, s *chat_service.ChatService) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{broker},
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 1e3, // 1KB
			MaxBytes: 1e6, // 1MB
		}),
		service: s,
	}
}

func (c *KafkaConsumer) Start() {
	log.Info().Msg("Kafka consumer start")
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Failed to read message")
			continue
		}
		log.Info().Msg("Consume message: " + string(m.Value))

		var message messageDto
		if err := json.Unmarshal(m.Value, &message); err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal message")
		} else {
			msg := models.Message{
				Content: message.Content,
				Role:    models.User,
			}
			go c.service.Ask(context.Background(), message.AuthToken, string(m.Key), &msg)
		}
	}
}
