package entities

// Numbers - https://wit.ai/docs/built-in-entities/20180601#wit_number_link
type Numbers struct {
	Items []Number `json:"number"`
}

// Number - represents the element of number
type Number struct {
	Confidence float64 `json:"confidence"`
	Value      int64   `json:"value"`
	Type       string  `json:"type"`
}
