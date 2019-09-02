package entities

// Quantities - https://wit.ai/docs/built-in-entities/20180601#wit_quantity_link
type Quantities struct {
	Items []Quantity `json:"quantity"`
}

// Quantity - represents the element of quantity
type Quantity struct {
	Confidence float64       `json:"confidence"`
	Value      *int64        `json:"value,omitempty"`
	Type       string        `json:"type"`
	From       *QuantityFrom `json:"from"`
	Product    *string       `json:"product,omitempty"`
	Unit       *string       `json:"unit,omitempty"`
}

// QuantityFrom - represents the from of quantity
type QuantityFrom struct {
	Value   string `json:"value"`
	Product string `json:"product"`
	Unit    string `json:"unit"`
}
