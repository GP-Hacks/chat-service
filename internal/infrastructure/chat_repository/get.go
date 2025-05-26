package chat_repository

import (
	"context"

	"github.com/GP-Hacks/chat/internal/models"
	"github.com/GP-Hacks/chat/internal/services"
)

func (r *ChatRepository) Get(ctx context.Context, userID, limit, offset int64) ([]models.Message, error) {
	q := `SELECT content, role FROM chat_history WHERE user_id = $1 ORDER BY created_at LIMIT $2 OFFSET $3`

	rows, err := r.pool.Query(ctx, q, userID, limit, offset)
	if err != nil {
		return nil, services.InternalServerError
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.Content, &msg.Role); err != nil {
			return nil, services.InternalServerError
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, services.InternalServerError
	}

	return messages, nil
}
