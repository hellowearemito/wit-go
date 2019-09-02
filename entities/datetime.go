package entities

// Datetimes - https://wit.ai/docs/built-in-entities/20180601#wit_datetime_link
type Datetimes struct {
	Items []Datetime `json:"datetime"`
}

// Datetime - represents the datetime element
type Datetime struct {
	Confidence float64        `json:"confidence"`
	Values     DatetimeValue  `json:"values"`
	Value      *string        `json:"value,omitempty"`
	Grain      *string        `json:"grain,omitempty"`
	To         *DatetimeRange `json:"to,omitempty"`
	From       *DatetimeRange `json:"from,omitempty"`
	Type       string         `json:"type"`
}

// DatetimeValue - represents the value of datetime
type DatetimeValue struct {
	Value *string        `json:"value,omitempty"`
	Grain *string        `json:"grain,omitempty"`
	To    *DatetimeRange `json:"to,omitempty"`
	From  *DatetimeRange `json:"from,omitempty"`
	Type  string         `json:"type"`
}

// DatetimeRange - represents the range of datetime
type DatetimeRange struct {
	Value string `json:"value"`
	Grain string `json:"grain"`
}
