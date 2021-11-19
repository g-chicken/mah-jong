package domain

import (
	"context"
	"fmt"
)

// HalfRoundGameScore is a score of half round game
// maximum of length is 4.
type HalfRoundGameScore []*PlayerScore

func (s HalfRoundGameScore) updateRanking() {
	indexOrderByScore := s.getIndexOrderByScore()

	for i, playerScoreIndex := range indexOrderByScore {
		s[playerScoreIndex].setRanking(uint32(i + 1))
	}
}

// UpdateScoresAndRankings updates scores and rankings of players.
// ranking is determined by scores.
// it is ignore if player ID in scores' keys are not player IDs of HandRoundGameScore.
func (s HalfRoundGameScore) updateScoresAndRankings(
	c context.Context, handID uint64, gameNumber uint32, scores map[uint64]int, /* [player ID] = score */
) error {
	if s == nil {
		return errNilHalfRoundGameScore
	}

	updatedHalfRoundGameScore := make(HalfRoundGameScore, 0, len(s))

	for _, playerScore := range s {
		score, ok := scores[playerScore.GetPlayerID()]
		if !ok {
			updatedHalfRoundGameScore = append(updatedHalfRoundGameScore, playerScore)

			continue
		}

		updatedHalfRoundGameScore = append(
			updatedHalfRoundGameScore,
			NewPlayerScore(playerScore.GetPlayerID(), score, playerScore.GetRanking()),
		)
	}

	updatedHalfRoundGameScore.updateRanking()

	if !updatedHalfRoundGameScore.Validate() {
		return fmt.Errorf("invalid score")
	}

	for i, playerScore := range s {
		if err := playerScore.updateScoreAndRankingByHandIDAndGameNumber(
			c, handID, updatedHalfRoundGameScore[i].GetScore(), updatedHalfRoundGameScore[i].GetRanking(), gameNumber,
		); err != nil {
			return err
		}
	}

	return nil
}

// Validate check HalfGameScore data.
// details of check are the following.
// - not the same player ID in []*PlayerScore.
// - []*PlayerScore's length is less or equal to 4.
// - sum of []*PlayerScore's score is zero.
// - correct ranking.
// NOTE: this method can optimize.
func (s HalfRoundGameScore) Validate() bool {
	return s.checkPlayers() && s.checkRanking() && s.checkSamePlayerID() && s.checkSumOfScores()
}

func (s HalfRoundGameScore) checkPlayers() bool {
	if len(s) > maxPlayersInHalfGame {
		return false
	}

	return true
}

func (s HalfRoundGameScore) checkSumOfScores() bool {
	sum := 0

	for _, playserScore := range s {
		sum += playserScore.GetScore()
	}

	return sum == 0
}

func (s HalfRoundGameScore) checkSamePlayerID() bool {
	playerIDs := make([]uint64, 0, len(s))

	for _, playerScore := range s {
		for _, playerID := range playerIDs {
			if playerID == playerScore.GetPlayerID() {
				return false
			}
		}

		playerIDs = append(playerIDs, playerScore.GetPlayerID())
	}

	return true
}

func (s HalfRoundGameScore) checkRanking() bool {
	indexOrderByScore := s.getIndexOrderByScore()

	for i, playerScoreIndex := range indexOrderByScore {
		if uint32(i+1) != s[playerScoreIndex].GetRanking() {
			return false
		}
	}

	return true
}

func (s HalfRoundGameScore) getIndexOrderByScore() []int {
	indexOrderByScore := make([]int, 0, len(s))

	for index := range s {
		indexOrderByScore = append(indexOrderByScore, index)
		for i := len(indexOrderByScore) - 1; i > 0; i-- {
			playerScoreIndex := indexOrderByScore[i]
			previousPlayerScoreIndex := indexOrderByScore[i-1]

			if s[previousPlayerScoreIndex].GetScore() < s[playerScoreIndex].GetScore() {
				indexOrderByScore[i], indexOrderByScore[i-1] = indexOrderByScore[i-1], indexOrderByScore[i]
			}
		}
	}

	return indexOrderByScore
}

// HalfGameScores is scores of half game.
// maximum of []*PlayerScore's length is 4.
type HalfRoundGameScores map[uint32]HalfRoundGameScore // key is the game number.

func (s HalfRoundGameScores) Validate() bool {
	for _, halfRoundGameScore := range s {
		if !halfRoundGameScore.Validate() {
			return false
		}
	}

	return true
}
