package servie_provider

import (
	"github.com/GP-Hacks/chat/internal/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ServiceProvider) AuthConnection() *grpc.ClientConn {
	if s.authConnection == nil {
		conn, err := grpc.Dial(
			config.Cfg.Grpc.AuthServiceAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		s.authConnection = conn
	}

	return s.authConnection
}
