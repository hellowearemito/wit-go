package entities

// PhraseToTranslates - https://wit.ai/docs/built-in-entities/20180601#wit_phrase_to_translate_link
type PhraseToTranslates struct {
	Items []PhraseToTranslate `json:"phrase_to_translate"`
}

// PhraseToTranslate - represents the element of phrase to translate
type PhraseToTranslate struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
