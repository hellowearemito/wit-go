package entities

// Emails - https://wit.ai/docs/built-in-entities/20180601#wit_email_link
type Emails struct {
	Items []Email `json:"email"`
}

// Email - represents email element.
type Email struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}
