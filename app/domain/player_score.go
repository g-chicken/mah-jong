package domain

import "context"

// PlayerScore is a score in a game.
type PlayerScore struct {
	playerID uint64
	score    int
	ranking  uint32
}

// NewPlayerScore create a PlayerScore.
func NewPlayerScore(playerID uint64, score int, ranking uint32) *PlayerScore {
	return &PlayerScore{
		playerID: playerID,
		score:    score,
		ranking:  ranking,
	}
}

// GetPlayerID returns the player ID.
func (p *PlayerScore) GetPlayerID() uint64 {
	if p == nil {
		return 0
	}

	return p.playerID
}

// GetScore returns the score.
func (p *PlayerScore) GetScore() int {
	if p == nil {
		return 0
	}

	return p.score
}

// GetRanking returns the ranking.
func (p *PlayerScore) GetRanking() uint32 {
	if p == nil {
		return 0
	}

	return p.ranking
}

func (p *PlayerScore) setScore(score int) {
	if p == nil {
		return
	}

	p.score = score
}

func (p *PlayerScore) setRanking(ranking uint32) {
	if p == nil {
		return
	}

	p.ranking = ranking
}

func (p *PlayerScore) updateScoreAndRankingByHandIDAndGameNumber(
	c context.Context, handID uint64, score int, ranking, gameNumber uint32,
) error {
	if p == nil {
		return errNilPlayerScore
	}

	if err := repos.halfRoundGameRepo.UpdateScoreAndRanking(c, handID, p.GetPlayerID(), score, ranking, gameNumber); err != nil {
		return err
	}

	p.setScore(score)
	p.setRanking(ranking)

	return nil
}
