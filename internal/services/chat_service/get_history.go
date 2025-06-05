package chat_service

import (
	"context"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/rs/zerolog/log"
)

func (s *ChatService) GetHistory(ctx context.Context, token string, limit, offset int64) ([]models.Message, error) {
	userID, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		log.Error().Msg("failed verify token: " + err.Error())
		return nil, err
	}

	his, err := s.chatRepository.Get(ctx, userID, limit, offset)
	if err != nil {
		log.Error().Msg("failed verify token: " + err.Error())
		return nil, err
	}

	return his, nil
}
