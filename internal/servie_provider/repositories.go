package servie_provider

import "github.com/GP-Hacks/chat/internal/infrastructure/chat_repository"

func (s *ServiceProvider) ChatRepository() *chat_repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chat_repository.NewChatRepository(s.DB())
	}

	return s.chatRepository
}
