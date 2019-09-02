package entities

// Durations - https://wit.ai/docs/built-in-entities/20180601#wit_duration_link
type Durations struct {
	Items []Duration `json:"duration"`
}

// Duration - represents the duration element.
type Duration struct {
	Confidence float64            `json:"confidence"`
	Second     int64              `json:"second"`
	Value      int64              `json:"value"`
	Type       string             `json:"type"`
	Unit       string             `json:"unit"`
	Normalized DurationNormalized `json:"normalized"`
}

// DurationNormalized - represents the normalized of duration element.
type DurationNormalized struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}
