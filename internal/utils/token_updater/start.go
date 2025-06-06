package tokenupdater

import (
	"fmt"
	"log"
)

func (tu *TokenUpdater) Start() error {
	tu.mu.Lock()
	defer tu.mu.Unlock()

	if tu.isRunning {
		return fmt.Errorf("token updater is already running")
	}

	if err := tu.updateToken(); err != nil {
		log.Printf("Initial token update failed: %v", err)
	}

	tu.isRunning = true

	go tu.run()

	log.Printf("Token updater started, will update every %v", tu.updateInterval)
	return nil
}
