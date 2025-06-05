package auth_adapter

import (
	"context"

	"github.com/GP-Hacks/chat/internal/services"
	"github.com/GP-Hacks/proto/pkg/api/auth"
	"github.com/rs/zerolog/log"
)

func (a *AuthAdapter) VerifyToken(ctx context.Context, token string) (int64, error) {
	resp, err := a.client.VerifyAccessToken(ctx, &auth.VerifyAccessTokenRequest{
		Access: token,
	})
	if err != nil {
		log.Debug().Msg(err.Error())
		return 0, services.InternalServerError
	}

	return resp.UserId, nil
}
