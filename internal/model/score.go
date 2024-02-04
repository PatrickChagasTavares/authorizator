package model

type Score struct {
	Punctuation int    `json:"pontuacao"`
	Source      string `json:"fonte" validate:"required"`
}

const minimumScore = 500

// ValidScore checks whether the score is higher than the minimum expected (500).
func (s *Score) ValidScore() bool {
	return s.Punctuation > minimumScore
}
