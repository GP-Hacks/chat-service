package chat_service

import (
	"context"

	"github.com/GP-Hacks/chat/internal/models"
)

type (
	IAuthAdapter interface {
		VerifyToken(ctx context.Context, token string) (int64, error)
	}

	IBotAdapter interface {
		Chat(ctx context.Context, messages ...models.Message) ([]models.Message, error)
	}

	IChatRepository interface {
		Add(ctx context.Context, userID int64, msg *models.Message) error
		Get(ctx context.Context, userID, limit, offset int64) ([]models.Message, error)
	}

	IBrokerProducer interface {
		Send(id string, msg *models.Message) error
		SendError(id string, err error) error
	}

	ChatService struct {
		authAdapter    IAuthAdapter
		botAdapter     IBotAdapter
		chatRepository IChatRepository
		brokerProducer IBrokerProducer
	}
)

func NewChatService(a IAuthAdapter, ba IBotAdapter, c IChatRepository, bp IBrokerProducer) *ChatService {
	return &ChatService{
		authAdapter:    a,
		botAdapter:     ba,
		chatRepository: c,
		brokerProducer: bp,
	}
}
