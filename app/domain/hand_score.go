package domain

import (
	"time"
)

// HandScore is a score of hand.
type HandScore struct {
	id                  uint32
	timestamp           time.Time
	halfRoundGameScores HalfRoundGameScores
}

// NewHandScore is a hand score model.
func NewHandScore(
	id uint32, timestamp time.Time, halfRoundGameScores HalfRoundGameScores,
) *HandScore {
	return &HandScore{
		id:                  id,
		timestamp:           timestamp,
		halfRoundGameScores: halfRoundGameScores,
	}
}

// GetID returns the hand ID.
func (h *HandScore) GetID() uint32 {
	if h == nil {
		return 0
	}

	return h.id
}

// GetTimestamp returns the timestamp.
func (h *HandScore) GetTimestamp() time.Time {
	if h == nil {
		return time.Time{}
	}

	return h.timestamp
}

// GetHalfGameScores returns the half game scores.
func (h *HandScore) GetHalfGameScores() HalfRoundGameScores {
	if h == nil {
		return HalfRoundGameScores{}
	}

	return h.halfRoundGameScores
}

// HalfGameScores is scores of half game.
// maximum of []*PlayerScore's length is 4.
type HalfRoundGameScores map[uint32][]*PlayerScore // key is the game number.

// Validate check HalfGameScores data.
// details of check are the following.
// - not the same player ID in []*PlayerScore.
// - []*PlayerScore's length is less or equal to 4.
// - sum of []*PlayerScore's score is zero.
// - correct ranking.
// NOTE: this method can optimize.
func (s HalfRoundGameScores) Validate() bool {
	return s.checkPlayers() && s.checkRanking() && s.checkSamePlayerID() && s.checkSumOfScores()
}

func (s HalfRoundGameScores) checkPlayers() bool {
	for _, playerScores := range s {
		if len(playerScores) > maxPlayersInHalfGame {
			return false
		}
	}

	return true
}

func (s HalfRoundGameScores) checkSumOfScores() bool {
	for _, playerScores := range s {
		sum := 0

		for _, playserScore := range playerScores {
			sum += playserScore.GetScore()
		}

		if sum != 0 {
			return false
		}
	}

	return true
}

func (s HalfRoundGameScores) checkSamePlayerID() bool {
	for _, playerScores := range s {
		playerIDs := make([]uint64, 0, len(playerScores))

		for _, playerScore := range playerScores {
			for _, playerID := range playerIDs {
				if playerID == playerScore.GetPlayerID() {
					return false
				}
			}

			playerIDs = append(playerIDs, playerScore.GetPlayerID())
		}
	}

	return true
}

func (s HalfRoundGameScores) checkRanking() bool {
	for _, playerScores := range s {
		indexOrderByScore := make([]int, 0, len(playerScores))

		for index := range playerScores {
			indexOrderByScore = append(indexOrderByScore, index)
			for i := len(indexOrderByScore) - 1; i > 0; i-- {
				playerScoreIndex := indexOrderByScore[i]
				previousPlayerScoreIndex := indexOrderByScore[i-1]

				if playerScores[previousPlayerScoreIndex].GetScore() < playerScores[playerScoreIndex].GetScore() {
					indexOrderByScore[i], indexOrderByScore[i-1] = indexOrderByScore[i-1], indexOrderByScore[i]
				}
			}
		}

		for i, playerScoreIndex := range indexOrderByScore {
			if uint32(i+1) != playerScores[playerScoreIndex].GetRanking() {
				return false
			}
		}
	}

	return true
}
