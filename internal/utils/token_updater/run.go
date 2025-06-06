package tokenupdater

import (
	"log"
	"time"
)

func (tu *TokenUpdater) run() {
	ticker := time.NewTicker(tu.updateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := tu.updateToken(); err != nil {
				log.Printf("Failed to update token: %v", err)
			}
		case <-tu.stopCh:
			return
		}
	}
}
