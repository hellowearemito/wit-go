package entities

// CreativeWorks - https://wit.ai/docs/built-in-entities/20180601#wit_creative_work_link
type CreativeWorks struct {
	Items []CreativeWork `json:"creative_work"`
}

// CreativeWork - represents the creative work element.
type CreativeWork struct {
	Confidence float64              `json:"confidence"`
	Value      string               `json:"value"`
	Resolved   CreativeWorkResolved `json:"resolved"`
}

// CreativeWorkResolved - represents the resolved of creative work.
type CreativeWorkResolved struct {
	Values []CreativeWorkResolve `json:"values"`
}

// CreativeWorkResolve - represents the resolve of creative work element.
type CreativeWorkResolve struct {
	Name     string   `json:"name"`
	Domain   string   `json:"domain"`
	Type     string   `json:"type"`
	External External `json:"external"`
}
