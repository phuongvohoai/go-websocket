package ws

type MessageEvent struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
}
