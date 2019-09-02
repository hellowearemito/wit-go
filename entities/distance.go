package entities

// Distances - https://wit.ai/docs/built-in-entities/20180601#wit_distance_link
type Distances struct {
	Items []Distance `json:"distance"`
}

// Distance - represents the distance element.
type Distance struct {
	Confidence float64        `json:"confidence"`
	Value      *int64         `json:"value,omitempty"`
	Type       string         `json:"type"`
	Unit       *string        `json:"unit,omitempty"`
	From       *DistanceRange `json:"from,omitempty"`
	To         *DistanceRange `json:"to,omitempty"`
}

// DistanceRange represents the from and to.
type DistanceRange struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}
