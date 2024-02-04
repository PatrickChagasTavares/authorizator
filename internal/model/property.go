package model

type Property struct {
	Value int `json:"valor" validate:"required"`
	Input int `json:"entrada" validate:"required"`
}

const (
	minimumPercent = 0.05 // 5%
	maximumPercent = 0.20 // 20%
)

// ValidMinInput
// Checks whether the input value is greater than or equal (>=) to the minimum expected value.
// The minimum expected value is 5% of the value of the property
func (p *Property) ValidMinInput() bool {
	minValue := int(float64(p.Value) * minimumPercent)

	return p.Input >= minValue
}

// ValidHighInput
// checks if the input value is less than the maximum expected value.
// The maximum expected value is 20%
func (p *Property) ValidHighInput() bool {
	maxValue := int(float64(p.Value) * maximumPercent)

	return p.Input < maxValue
}
