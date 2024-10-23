package chat

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistent = "assistent"
)

type Message struct {
	Role    string `json:"role" db:"role"`
	Content string `json:"content" db:"content"`
}

type HistoryRequest struct {
	UserId int `json:"user_id"`
	ChatId int `json:"chat_id"`
}

type HistoryResponse struct {
	UserId   int       `json:"user_id"`
	ChatId   int       `json:"chat_id"`
	Messages []Message `json:"messages"`
}

type ChatItem struct {
	UserId  int `json:"user_id"`
	ChatId  int `json:"chat_id"`
	Message `json:"message"`
}

type ChatPreview struct {
	ChatId  int    `json:"chat_id"`
	Content string `json:"content"`
}

type StartResponse struct {
	UserId int           `json:"user_id"`
	Chats  []ChatPreview `json:"chats"`
}
