package tokenupdater

import (
	"fmt"
	"log"
	"time"
)

func (tu *TokenUpdater) updateToken() error {
	tokenResp, err := tu.fetchNewToken()
	if err != nil {
		return fmt.Errorf("failed to fetch new token: %w", err)
	}

	tu.CurrentToken = tokenResp
	log.Printf("Token updated successfully at %v", time.Now().Format(time.RFC3339))

	return nil
}
