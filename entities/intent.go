package entities

// Intents - represents the list of intent
type Intents struct {
	Items []Intent `json:"intent"`
}

// Intent - represents the intent element
type Intent struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
