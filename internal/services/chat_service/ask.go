package chat_service

import (
	"context"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/GP-Hacks/chat/internal/services"
)

func (s *ChatService) Ask(ctx context.Context, token string, id string, msg *models.Message) {
	userID, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		s.brokerProducer.SendError(id, err)
		return
	}

	history, err := s.chatRepository.Get(ctx, userID, 5, 0)
	if err != nil {
		s.brokerProducer.SendError(id, err)
		return
	}

	history = append(history, *msg)

	asw, err := s.botAdapter.Chat(ctx, history...)
	if len(asw)+1 != len(history) {
		s.brokerProducer.SendError(id, services.InternalServerError)
		return
	}

	if err := s.chatRepository.Add(ctx, userID, &asw[len(asw)-1]); err != nil {
		s.brokerProducer.SendError(id, err)
		return
	}

	if err := s.brokerProducer.Send(id, &asw[len(asw)-1]); err != nil {

	}
}
