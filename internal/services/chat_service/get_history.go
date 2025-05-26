package chat_service

import (
	"context"

	"github.com/GP-Hacks/chat/internal/models"
)

func (s *ChatService) GetHistory(ctx context.Context, token string, limit, offset int64) ([]models.Message, error) {
	userID, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return s.chatRepository.Get(ctx, userID, limit, offset)
}
