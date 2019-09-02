package entities

// LocalSearchQueries - https://wit.ai/docs/built-in-entities/20180601#wit_local_search_query_link
type LocalSearchQueries struct {
	Items []LocalSearchQuery `json:"local_search_query"`
}

// LocalSearchQuery - represents the local search query element
type LocalSearchQuery struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
