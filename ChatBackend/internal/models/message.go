package models

const (
	RoleSystem    string = "system"
	RoleUser      string = "user"
	RoleAssistant string = "assistant"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
