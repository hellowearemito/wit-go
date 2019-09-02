package entities

// Locations - https://wit.ai/docs/built-in-entities/20180601#wit_location_link
type Locations struct {
	Items []Location `json:"location"`
}

// Location - represents the element of location
type Location struct {
	Confidence float64           `json:"confidence"`
	Value      string            `json:"value"`
	Resolved   *LocationResolved `json:"resolved,omitempty"`
	Type       *string           `json:"type,omitempty"`
}

// LocationResolved - represents the resolved of location
type LocationResolved struct {
	Values []LocationResolvedValue `json:"values"`
}

// LocationResolvedValue - represents the value of location resolved
type LocationResolvedValue struct {
	Name     string           `json:"name"`
	Grain    string           `json:"grain"`
	Type     string           `json:"type"`
	Timezone string           `json:"timezone"`
	Coords   LocationCoords   `json:"coords"`
	External LocationExternal `json:"external"`
}

// LocationCoords - represents the coords of location
type LocationCoords struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

// LocationExternal - represents the external of location
type LocationExternal struct {
	Geonames  string `json:"geonames"`
	Wikipedia string `json:"wikipedia"`
}
