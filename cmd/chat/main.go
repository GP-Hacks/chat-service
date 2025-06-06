package main

import (
	"net"

	"github.com/GP-Hacks/chat/internal/config"
	"github.com/GP-Hacks/chat/internal/servie_provider"
	"github.com/GP-Hacks/chat/internal/utils/logger"
	proto "github.com/GP-Hacks/proto/pkg/api/chat"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig("./config")
	logger.SetupLogger()
	serviceProvider := servie_provider.NewServiceProvider()

	log.Info().Msg("Init app")

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterChatServiceServer(grpcServer, serviceProvider.ChatController())

	go serviceProvider.KafkaConsumer().Start()

	list, err := net.Listen("tcp", ":"+config.Cfg.Grpc.Port)
	if err != nil {
		log.Fatal().Msg("Failed start listen port")
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal().Msg("Failed serve grpc")
	}
}
