package broker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

type respDto struct {
	Status    string    `json:"status,omitempty"`
	Error     string    `json:"error,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewKafkaProducer(broker, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaProducer) Send(id string, msg *models.Message) error {
	dto := respDto{
		Status:    "success",
		Content:   msg.Content,
		CreatedAt: time.Now(), //TODO: replace to real rime
	}
	data, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		err = p.writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(id),
			Value: data,
		})

		if err == nil {
			return nil
		}
	}

	return err
}

func (p *KafkaProducer) SendError(id string, err error) error {
	dto := respDto{
		Status: "error",
		Error:  err.Error(),
	}
	data, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		err = p.writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(id),
			Value: data,
		})

		if err == nil {
			return nil
		}
	}

	return err
}
