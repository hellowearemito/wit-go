package entities

// Volumes - https://wit.ai/docs/built-in-entities/20180601#wit_volume_link
type Volumes struct {
	Items []Volume `json:"volume"`
}

// Volume - represents the element of volume
type Volume struct {
	Confidence float64 `json:"confidence"`
	Value      int64   `json:"value"`
	Type       string  `json:"type"`
	Unit       string  `json:"unit"`
}
