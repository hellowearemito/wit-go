package entities

// Contacts - https://wit.ai/docs/built-in-entities/20180601#wit_contact_link
type Contacts struct {
	Items []Contact `json:"contact"`
}

// Contact - represents the contact element.
type Contact struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
