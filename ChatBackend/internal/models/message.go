package models

const (
	RoleSystem    string = "system"
	RoleUser      string = "user"
	RoleAssistant string = "assistant"
)

type ContentType int

const (
	EmptyType ContentType = 0
	TextType  ContentType = 1
	AudioType ContentType = 2
	VideoType ContentType = 3
)

const (
	EmptyTypePlug = "Новый чат"
	AudioTypePlug = "**Аудио**"
	VideoTypePlug = "**Видео**"
)

type Message struct {
	Role        string      `json:"role" db:"role"`
	Content     string      `json:"content" db:"content"`
	ContentType ContentType `json:"content_type" db:"content_type"`
}
