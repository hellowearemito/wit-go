package entities

// AgendaEntries - https://wit.ai/docs/built-in-entities/20180601#wit_agenda_entry_link
type AgendaEntries struct {
	Items []AgendaEntries `json:"agenda_entry"`
}

// AgendaEntry - represents the agenda entry element.
type AgendaEntry struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
