package model

type Person struct {
	Name     string `json:"nome" validate:"required"`
	Document string `json:"cpf" validate:"required"`
	Income   int    `json:"renda"`
}

const minimumIncome = 2000

// ValidIncome checks whether the person's income is greater than the minimum expected (20000).
func (p *Person) ValidIncome() bool {
	return p.Income > minimumIncome
}
