package entities

// Urls - https://wit.ai/docs/built-in-entities/20180601#wit_url_link
type Urls struct {
	Items []URL `json:"url"`
}

// URL represents the element of url
type URL struct {
	Confidence float64 `json:"confidence"`
	Domain     string  `json:"domain"`
	Value      string  `json:"value"`
}
