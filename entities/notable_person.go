package entities

// NotablePeople - https://wit.ai/docs/built-in-entities/20180601#wit_notable_person_link
type NotablePeople struct {
	Items []NotablePerson `json:"notable_person"`
}

// NotablePerson - represents the element of person of notable
type NotablePerson struct {
	Confidence float64            `json:"confidence"`
	Value      NotablePersonValue `json:"value"`
	Type       string             `json:"type"`
}

// NotablePersonValue - represents the value of person of notable
type NotablePersonValue struct {
	Name     string   `json:"name"`
	External External `json:"external"`
}
