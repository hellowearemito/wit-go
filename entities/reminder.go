package entities

// Reminders - https://wit.ai/docs/built-in-entities/20180601#wit_reminder_link
type Reminders struct {
	Items []Reminder `json:"reminder"`
}

// Reminder - represents the element of reminder
type Reminder struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
