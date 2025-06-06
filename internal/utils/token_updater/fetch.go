package tokenupdater

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GP-Hacks/chat/internal/config"
)

type tokenResponse struct {
	Token     string    `json:"iamToken"`
	ExpiresIn time.Time `json:"expiresAt"`
}

type tokenRequest struct {
	OAuthToken string `json:"yandexPassportOauthToken"`
}

func (tu *TokenUpdater) fetchNewToken() (string, error) {
	req := tokenRequest{
		OAuthToken: config.Cfg.AIModel.AuthToken,
	}
	reqM, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed marshal req")
	}

	r := bytes.NewReader(reqM)
	resp, err := http.Post(tu.tokenURL, "application/json", r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var tokenResp tokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if tokenResp.Token == "" {
		return "", fmt.Errorf("empty token in response")
	}

	return tokenResp.Token, nil
}
