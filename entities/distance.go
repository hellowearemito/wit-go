package entities

// Distances - https://wit.ai/docs/built-in-entities/20180601#wit_distance_link
type Distances struct {
	Items []Distance `json:"distance"`
}

// Distance - represents the distance element.
type Distance struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
	Unit       string  `json:"unit"`
}
