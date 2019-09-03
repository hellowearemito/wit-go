package entities

// Entity - represents the entity element
type Entity struct {
	Confidence float64   `json:"confidence"`
	Entities   *Entities `json:"entities,omitempty"`
	Value      string    `json:"value"`
	Type       string    `json:"type"`
	Suggested  bool      `json:"suggested"`
}
