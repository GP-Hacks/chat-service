package servie_provider

import (
	tokenupdater "github.com/GP-Hacks/chat/internal/utils/token_updater"
)

func (s *ServiceProvider) TokenUpdater() *tokenupdater.TokenUpdater {
	if s.tokenUpdater == nil {
		s.tokenUpdater = tokenupdater.NewTokenUpdater("https://iam.api.cloud.yandex.net/iam/v1/tokens")
	}

	return s.tokenUpdater
}
