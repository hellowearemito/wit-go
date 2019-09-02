package entities

// Ordinals - https://wit.ai/docs/built-in-entities/20180601#wit_ordinal_link
type Ordinals struct {
	Items []Ordinal `json:"ordinal"`
}

// Ordinal - represents the element of ordinal
type Ordinal struct {
	Confidence float64 `json:"confidence"`
	Value      int64   `json:"value"`
	Type       string  `json:"type"`
}
