package tokenupdater

import "log"

func (tu *TokenUpdater) Stop() {
	tu.mu.Lock()
	defer tu.mu.Unlock()

	if !tu.isRunning {
		return
	}

	close(tu.stopCh)
	tu.isRunning = false
	log.Println("Token updater stopped")
}
