package models

type SenderType string

const (
	Bot  SenderType = "bot"
	User SenderType = "user"
)

type Message struct {
	Content string
	Role    SenderType
}
