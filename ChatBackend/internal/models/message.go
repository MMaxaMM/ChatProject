package models

const (
	RoleSystem    string = "system"
	RoleUser      string = "user"
	RoleAssistant string = "assistant"
)

type ContentType int

const (
	Empty     ContentType = 0
	TextType  ContentType = 1
	AudioType ContentType = 2
	VideoType ContentType = 3
)

type Message struct {
	Role        string      `json:"role"`
	Content     string      `json:"content"`
	ContentType ContentType `json:"content_type"`
}
