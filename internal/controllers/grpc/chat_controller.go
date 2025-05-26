package grpc

import (
	"context"

	desc "github.com/GP-Hacks/proto/pkg/api/chat"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/GP-Hacks/chat/internal/services/chat_service"
)

type ChatController struct {
	desc.UnimplementedChatServiceServer
	chatService *chat_service.ChatService
}

func NewChatController(cs *chat_service.ChatService) *ChatController {
	return &ChatController{
		chatService: cs,
	}
}

func (c *ChatController) GetHistory(ctx context.Context, req *desc.GetHistoryRequest) (*desc.GetHistoryResponse, error) {
	his, err := c.chatService.GetHistory(ctx, req.Token, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	res := make([]*desc.ChatMessage, len(his))
	for i, m := range his {
		var r desc.ChatRole
		if m.Role == models.Bot {
			r = desc.ChatRole_BOT
		} else {
			r = desc.ChatRole_USER
		}

		res[i] = &desc.ChatMessage{
			Content:   m.Content,
			Role:      r,
			CreatedAt: timestamppb.Now(),
		}
	}

	return &desc.GetHistoryResponse{
		Messages: res,
	}, nil
}
