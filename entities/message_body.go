package entities

// MessageBodies - https://wit.ai/docs/built-in-entities/20180601#wit_message_body_link
type MessageBodies struct {
	Items []MessageBody `json:"message_body"`
}

// MessageBody - represents the element of body of message
type MessageBody struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
