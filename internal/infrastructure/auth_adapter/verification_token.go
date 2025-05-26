package auth_adapter

import (
	"context"

	"github.com/GP-Hacks/chat/internal/services"
	"github.com/GP-Hacks/proto/pkg/api/auth"
)

func (a *AuthAdapter) VerifyToken(ctx context.Context, token string) (int64, error) {
	resp, err := a.client.VerifyAccessToken(ctx, &auth.VerifyAccessTokenRequest{
		Access: token,
	})

	if err != nil {
		return 0, services.InternalServerError
	}

	return resp.UserId, nil
}
