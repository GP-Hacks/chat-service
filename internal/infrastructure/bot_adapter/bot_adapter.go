package bot_adapter

import tokenupdater "github.com/GP-Hacks/chat/internal/utils/token_updater"

type BotAdapter struct {
	context      string
	tokenUpdater *tokenupdater.TokenUpdater
}

func NewBotAdapter(context string, tokenupdater *tokenupdater.TokenUpdater) *BotAdapter {
	return &BotAdapter{
		context:      context,
		tokenUpdater: tokenupdater,
	}
}
