package servie_provider

import "github.com/GP-Hacks/chat/internal/services/chat_service"

func (s *ServiceProvider) ChatService() *chat_service.ChatService {
	if s.chatService == nil {
		s.chatService = chat_service.NewChatService(s.AuthAdapter(), s.BotAdapter(), s.ChatRepository(), s.KafkaProducer())
	}

	return s.chatService
}
