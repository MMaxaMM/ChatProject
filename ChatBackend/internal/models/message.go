package models

const (
	RoleSystem    string = "system"
	RoleUser      string = "user"
	RoleAssistant string = "assistant"
)

const (
	Empty     int = 0
	TextType  int = 1
	AudioType int = 2
	VideoType int = 3
)

type Message struct {
	Role        string `json:"role" db:"role"`
	Content     string `json:"content" db:"content"`
	ContentType int    `json:"content_type" db:"content_type"`
}
