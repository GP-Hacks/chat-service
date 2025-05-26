package auth_adapter

import (
	desc "github.com/GP-Hacks/proto/pkg/api/auth"
)

type AuthAdapter struct {
	client desc.AuthServiceClient
}

func NewAuthAdapter(c desc.AuthServiceClient) *AuthAdapter {
	return &AuthAdapter{
		client: c,
	}
}
