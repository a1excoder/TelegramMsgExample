package TgTypes

type Message struct {
	MessageId int64            `json:"message_id"`
	From      *User            `json:"from,omitempty"`
	Chat      *Chat            `json:"chat,omitempty"`
	Date      int              `json:"date"`
	Text      *string          `json:"text,omitempty"`
	Entities  *[]MessageEntity `json:"entities,omitempty"`
}

type Chat struct {
	Id        int64   `json:"id"`
	Type      string  `json:"type"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Username  *string `json:"username,omitempty"`
}

type User struct {
	Id           int64   `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name,omitempty"`
	Username     *string `json:"username,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"`
}

type MessageEntity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}
