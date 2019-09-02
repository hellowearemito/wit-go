package entities

// WolframSearchQueries - https://wit.ai/docs/built-in-entities/20180601#wit_wolfram_search_query_link
type WolframSearchQueries struct {
	Items []WolframSearchQuery `json:"wolfram_search_query"`
}

// WolframSearchQuery - represents the element of wolframe search query
type WolframSearchQuery struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
