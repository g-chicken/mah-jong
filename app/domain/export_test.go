package domain

import "context"

/**
 * HalfRoundGameScore
 *.*/
func (s HalfRoundGameScore) UpdateScoreAndRankings(
	c context.Context, handID uint64, gameNumber uint32, scores map[uint64]int,
) error {
	return s.updateScoresAndRankings(c, handID, gameNumber, scores)
}

func (s HalfRoundGameScore) UpdateRanking() { s.updateRanking() }

func (s HalfRoundGameScore) Validate() bool {
	return s.validate()
}
