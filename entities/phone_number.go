package entities

// PhoneNumbers - https://wit.ai/docs/built-in-entities/20180601#wit_phone_number_link
type PhoneNumbers struct {
	Items []PhoneNumber `json:"phone_number"`
}

// PhoneNumber - represents element of phone number
type PhoneNumber struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
