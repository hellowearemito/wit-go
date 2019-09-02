package entities

// MathExpressions - https://wit.ai/docs/built-in-entities/20180601#wit_math_expression_link
type MathExpressions struct {
	Items []MathExpression `json:"math_expression"`
}

// MathExpression - represents the element of expression of math
type MathExpression struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
	Type       string  `json:"type"`
}
