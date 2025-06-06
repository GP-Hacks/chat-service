package tokenupdater

import (
	"net/http"
	"sync"
	"time"
)

type TokenUpdater struct {
	tokenURL       string
	updateInterval time.Duration
	httpClient     *http.Client
	stopCh         chan struct{}
	mu             sync.Mutex
	isRunning      bool

	CurrentToken string
}

func NewTokenUpdater(tokenURL string) *TokenUpdater {
	return &TokenUpdater{
		tokenURL:       tokenURL,
		updateInterval: 10 * time.Hour,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		stopCh: make(chan struct{}),
	}
}
