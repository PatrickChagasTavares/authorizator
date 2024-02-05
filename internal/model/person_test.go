package model

func (suite *ModelSuite) TestValidIncomeSuccess() {
	person := Person{Name: "John Doe", Document: "12345678901", Income: 2500}

	isValid := person.ValidIncome()

	suite.True(isValid)
}

func (suite *ModelSuite) TestValidIncomeError() {
	person := Person{Name: "Jane Doe", Document: "98765432101", Income: 1500}

	isValid := person.ValidIncome()

	suite.False(isValid)
}
