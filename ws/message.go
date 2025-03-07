package ws

import "encoding/json"

type MessageEvent struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
}

func NewMessageEvent(user, text string) MessageEvent {
	return MessageEvent{
		Type: "message",
		User: user,
		Text: text,
	}
}

func (m *MessageEvent) ToJson() []byte {
	jsonString, _ := json.Marshal(m)
	return jsonString
}
