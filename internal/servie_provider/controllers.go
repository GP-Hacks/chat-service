package servie_provider

import "github.com/GP-Hacks/chat/internal/controllers/grpc"

func (s *ServiceProvider) ChatController() *grpc.ChatController {
	if s.chatController == nil {
		s.chatController = grpc.NewChatController(s.ChatService())
	}

	return s.chatController
}
