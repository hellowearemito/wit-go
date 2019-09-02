package entities

// WikipediaSearchQueries - https://wit.ai/docs/built-in-entities/20180601#wit_wikipedia_search_query_link
type WikipediaSearchQueries struct {
	Items []WikipediaSearchQuery `json:"wikipedia_search_query"`
}

// WikipediaSearchQuery - represents the element of wikipedia search query
type WikipediaSearchQuery struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
