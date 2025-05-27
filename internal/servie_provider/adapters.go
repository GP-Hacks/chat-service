package servie_provider

import (
	"github.com/GP-Hacks/chat/internal/config"
	"github.com/GP-Hacks/chat/internal/infrastructure/auth_adapter"
	"github.com/GP-Hacks/chat/internal/infrastructure/bot_adapter"
)

func (s *ServiceProvider) AuthAdapter() *auth_adapter.AuthAdapter {
	if s.authAdapter == nil {
		s.authAdapter = auth_adapter.NewAuthAdapter(s.AuthClient())
	}

	return s.authAdapter
}

func (s *ServiceProvider) BotAdapter() *bot_adapter.BotAdapter {
	if s.botAdapter == nil {
		s.botAdapter = bot_adapter.NewBotAdapter(config.Cfg.AIModel.BaseContext)
	}

	return s.botAdapter
}
