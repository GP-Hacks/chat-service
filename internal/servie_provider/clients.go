package servie_provider

import "github.com/GP-Hacks/proto/pkg/api/auth"

func (s *ServiceProvider) AuthClient() auth.AuthServiceClient {
	if s.authClient == nil {
		s.authClient = auth.NewAuthServiceClient(s.AuthConnection())
	}

	return s.authClient
}
