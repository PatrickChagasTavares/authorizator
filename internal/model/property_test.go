package model

func (suite *ModelSuite) TestValidMinInputSuccess() {
	property := Property{Value: 100000, Input: 6000}

	isValid := property.ValidMinInput()

	suite.True(isValid)
}

func (suite *ModelSuite) TestValidMinInputError() {
	property := Property{Value: 200000, Input: 8000}

	isValid := property.ValidMinInput()

	suite.False(isValid)
}

func (suite *ModelSuite) TestValidHighInputSuccess() {
	property := Property{Value: 150000, Input: 25000}

	isValid := property.ValidHighInput()

	suite.True(isValid)
}

func (suite *ModelSuite) TestValidHighInputError() {
	property := Property{Value: 300000, Input: 70000}

	isValid := property.ValidHighInput()

	suite.False(isValid)
}
