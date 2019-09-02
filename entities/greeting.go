package entities

// Byes - https://wit.ai/docs/built-in-entities/20180601#wit_bye_link
type Byes struct {
	Items []Bye `json:"bye"`
}

// Bye - represents the bye element
type Bye struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}

// Greetings - https://wit.ai/docs/built-in-entities/20180601#wit_greetings_link
type Greetings struct {
	Items []Greeting `json:"greetings"`
}

// Greeting - represents the greeting element
type Greeting struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
