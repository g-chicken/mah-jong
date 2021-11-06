package domain

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
