package entities

// Temperatures - https://wit.ai/docs/built-in-entities/20180601#wit_temperature_link
type Temperatures struct {
	Items []Temperature `json:"temperature"`
}

// Temperature - represents the element of temperature
type Temperature struct {
	Confidence float64          `json:"confidence"`
	Type       string           `json:"type"`
	From       *TemperatureFrom `json:"from"`
	Value      *int64           `json:"value,omitempty"`
	Unit       *string          `json:"unit"`
}

// TemperatureFrom - represents the from of temperature
type TemperatureFrom struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}
