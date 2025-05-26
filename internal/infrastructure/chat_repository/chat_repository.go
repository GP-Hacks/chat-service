package chat_repository

import "github.com/jackc/pgx/v5/pgxpool"

type ChatRepository struct {
	pool *pgxpool.Pool
}

func NewChatRepository(p *pgxpool.Pool) *ChatRepository {
	return &ChatRepository{
		pool: p,
	}
}
