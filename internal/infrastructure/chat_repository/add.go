package chat_repository

import (
	"context"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/GP-Hacks/chat/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func (r *ChatRepository) Add(ctx context.Context, userID int64, msg *models.Message) error {
	q := `INSERT INTO chat_history (user_id, content, role) VALUES ($1, $2, $3, $4, $5)`

	if _, err := r.pool.Exec(ctx, q, userID, msg.Content, msg.Role); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return services.AlreadyExistsError
			}
		}

		log.Error().Msg(err.Error())
		return services.InternalServerError
	}

	return nil
}
