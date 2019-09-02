package entities

// Sentiments - https://wit.ai/docs/built-in-entities/20180601#wit_sentiment_link
type Sentiments struct {
	Items []Sentiment `json:"items"`
}

// Sentiment - represents the element of sentiment
type Sentiment struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
