package entities

// AmountOfMonies - https://wit.ai/docs/built-in-entities/20180601#wit_amount_of_money_link
type AmountOfMonies struct {
	Items []AmountOfMoney `json:"amount_of_money"`
}

// AmountOfMoney - represents the amount of money element.
type AmountOfMoney struct {
	Confidence float64             `json:"confidence"`
	From       *AmountOfMoneyRange `json:"from,omitempty"`
	To         *AmountOfMoneyRange `json:"to,omitempty"`
	Type       string              `json:"type"`
	Unit       *string             `json:"unit,omitempty"`
}

// AmountOfMoney - represents the range of amount of money element.
type AmountOfMoneyRange struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}
