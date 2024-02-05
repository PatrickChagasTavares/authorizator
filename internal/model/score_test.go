package model

func (suite *ModelSuite) TestValidScoreSuccess() {
	score := Score{Punctuation: 600, Source: "serasa"}

	isValid := score.ValidScore()

	suite.True(isValid)
}

func (suite *ModelSuite) TestValidScoreError() {
	score := Score{Punctuation: 400, Source: "serasa"}

	isValid := score.ValidScore()

	suite.False(isValid)
}
