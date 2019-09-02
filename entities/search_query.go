package entities

// SearchQueries - https://wit.ai/docs/built-in-entities/20180601#wit_search_query_link
type SearchQueries struct {
	Items []SearchQuery `json:"search_query"`
}

// SearchQuery - represents the element of search query
type SearchQuery struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
