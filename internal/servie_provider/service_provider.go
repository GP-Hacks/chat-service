package servie_provider

import (
	controllers "github.com/GP-Hacks/chat/internal/controllers/grpc"
	"github.com/GP-Hacks/chat/internal/infrastructure/auth_adapter"
	"github.com/GP-Hacks/chat/internal/infrastructure/bot_adapter"
	"github.com/GP-Hacks/chat/internal/infrastructure/broker"
	"github.com/GP-Hacks/chat/internal/infrastructure/chat_repository"
	"github.com/GP-Hacks/chat/internal/services/chat_service"
	tokenupdater "github.com/GP-Hacks/chat/internal/utils/token_updater"
	"github.com/GP-Hacks/proto/pkg/api/auth"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type ServiceProvider struct {
	db             *pgxpool.Pool
	authConnection *grpc.ClientConn
	authClient     auth.AuthServiceClient
	tokenUpdater   *tokenupdater.TokenUpdater

	kafkaConsumer  *broker.KafkaConsumer
	kafkaProducer  *broker.KafkaProducer
	authAdapter    *auth_adapter.AuthAdapter
	botAdapter     *bot_adapter.BotAdapter
	chatRepository *chat_repository.ChatRepository

	chatService *chat_service.ChatService

	chatController *controllers.ChatController
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
