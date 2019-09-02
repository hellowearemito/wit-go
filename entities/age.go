package entities

// AgeOfPersons - https://wit.ai/docs/built-in-entities/20180601#wit_age_of_person_link
type AgeOfPersons struct {
	Items []AgeOfPerson `json:"age_of_person"`
}

// AgeOfPerson - represents the age of person element.
type AgeOfPerson struct {
	Value      string  `json:"value"`
	Type       string  `json:"type"`
	Confidence float64 `json:"confidence"`
}
