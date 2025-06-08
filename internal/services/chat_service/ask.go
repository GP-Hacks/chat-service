package chat_service

import (
	"context"
	"fmt"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/GP-Hacks/chat/internal/services"
	"github.com/rs/zerolog/log"
)

func (s *ChatService) Ask(ctx context.Context, token string, id string, msg *models.Message) {
	userID, err := s.authAdapter.VerifyToken(ctx, token)
	if err != nil {
		log.Error().Msg("failed verify token: " + err.Error())
		s.brokerProducer.SendError(id, err)
		return
	}

	history, err := s.chatRepository.Get(ctx, userID, 5, 0)
	if err != nil {
		log.Error().Msg("failed get history: " + err.Error())
		s.brokerProducer.SendError(id, err)
		return
	}

	if err := s.chatRepository.Add(ctx, userID, msg); err != nil {
		log.Error().Msg("failed question answer: " + err.Error())
		s.brokerProducer.SendError(id, err)
		return
	}

	history = append(history, *msg)

	asw, err := s.botAdapter.Chat(ctx, history...)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed get answer: %v", err))
		if err := s.chatRepository.Add(ctx, userID, &models.Message{Content: "internal server error", Role: models.Bot}); err != nil {
			log.Error().Msg("failed question answer: " + err.Error())
			s.brokerProducer.SendError(id, err)
			return
		}

		s.brokerProducer.SendError(id, services.InternalServerError)
		return
	}

	if err := s.chatRepository.Add(ctx, userID, &asw[len(asw)-1]); err != nil {
		log.Error().Msg("failed write answer: " + err.Error())
		if err := s.chatRepository.Add(ctx, userID, &models.Message{Content: "internal server error", Role: models.Bot}); err != nil {
			log.Error().Msg("failed question answer: " + err.Error())
			s.brokerProducer.SendError(id, err)
			return
		}
		s.brokerProducer.SendError(id, err)
		return
	}

	if err := s.brokerProducer.Send(id, &asw[len(asw)-1]); err != nil {
		if err := s.chatRepository.Add(ctx, userID, &models.Message{Content: "internal server error", Role: models.Bot}); err != nil {
			log.Error().Msg("failed question answer: " + err.Error())
			s.brokerProducer.SendError(id, err)
			return
		}
		log.Error().Msg("failed send answer to broker: " + err.Error())
	}
}
