package entities

// Thanks - https://wit.ai/docs/built-in-entities/20180601#wit_thanks_link
type Thanks struct {
	Items []Thank `json:"thanks"`
}

// Thank - represents the element of thank
type Thank struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
