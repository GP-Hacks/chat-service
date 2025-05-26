package bot_adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/GP-Hacks/chat/internal/config"
	"github.com/GP-Hacks/chat/internal/models"
)

type (
	reqDto struct {
		ModelUri          string               `json:"model_uri"`
		CompletionOptions completionOptionsDto `json:"completion_options"`
		Messages          []messageDto         `json:"messages"`
	}

	messageDto struct {
		Role string `json:"role"`
		Text string `json:"text"`
	}

	completionOptionsDto struct {
		MaxTokens   int     `json:"max_tokens"`
		Temperature float32 `json:"temperature"`
	}

	responseDto struct {
		Result resultDto `json:"result"`
	}

	resultDto struct {
		Alternatives []alternativeDto `json:"alternatives"`
	}

	alternativeDto struct {
		Message messageDto `json:"message"`
		Status  string     `json:"status"`
	}
)

func convertMessagesToDto(messages ...models.Message) []messageDto {
	res := make([]messageDto, 0, len(messages))
	for _, el := range messages {
		if el.Role == models.User {
			res = append(res, messageDto{
				Role: "user",
				Text: el.Content,
			})
		} else {
			res = append(res, messageDto{
				Role: "assistant",
				Text: el.Content,
			})
		}
	}

	return res
}

func convertDtoToMessages(messagesDto ...alternativeDto) []models.Message {
	res := make([]models.Message, 0, len(messagesDto))
	for _, el := range messagesDto {
		if el.Message.Role == "user" {
			res = append(res, models.Message{
				Role:    models.User,
				Content: el.Message.Text,
			})
		} else if el.Message.Role == "assistant" {
			res = append(res, models.Message{
				Role:    models.Bot,
				Content: el.Message.Text,
			})
		}
	}

	return res
}

func (r *BotAdapter) Chat(ctx context.Context, messages ...models.Message) ([]models.Message, error) {
	messagesDto := make([]messageDto, 0, len(messages)+1)
	messagesDto = append(messagesDto, messageDto{
		Role: "system",
		Text: r.context,
	})
	messagesDto = append(messagesDto, convertMessagesToDto(messages...)...)
	optDto := completionOptionsDto{
		MaxTokens:   config.Cfg.AIModel.MaxTokens,
		Temperature: config.Cfg.AIModel.Temperature,
	}
	reqDto := reqDto{
		ModelUri:          config.Cfg.AIModel.ModelUri,
		CompletionOptions: optDto,
		Messages:          messagesDto,
	}

	jsonData, err := json.Marshal(reqDto)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.Cfg.AIModel.Address, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.AIModel.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseDTO responseDto
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	return convertDtoToMessages(responseDTO.Result.Alternatives...), nil
}
