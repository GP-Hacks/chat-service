package bot_adapter

type BotAdapter struct {
	context string
}

func NewBotAdapter(context string) *BotAdapter {
	return &BotAdapter{
		context: context,
	}
}
