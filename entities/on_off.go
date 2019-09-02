package entities

// OnOffs - https://wit.ai/docs/built-in-entities/20180601#wit_on_off_link
type OnOffs struct {
	Items []OnOff `json:"on_off"`
}

// OnOff - represents the element of on-off
type OnOff struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
