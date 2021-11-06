package rdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/g-chicken/mah-jong/app/domain"
)

type halfRoundGameRepository struct {
	repo domain.RDBDetectorRepository
}

// NewHalfRoundGameRepository implements the domain.HalfRoundGameRepository.
func NewHalfRoundGameRepository(repo domain.RDBDetectorRepository) domain.HalfRoundGameRepository {
	return &halfRoundGameRepository{
		repo: repo,
	}
}

func (r *halfRoundGameRepository) CreateHalfRoundGames(
	c context.Context, handID uint64, halfRoundGameScores domain.HalfRoundGameScores,
) error {
	if len(halfRoundGameScores) == 0 {
		return nil
	}

	// may not hove to use validate method...
	if !halfRoundGameScores.Validate() {
		return domain.NewInvalidArgumentError("invalid argument of CreateHalfRoundGames")
	}

	ope := r.repo.GetRDBOperator(c)

	argsStatement := []string{}
	args := []interface{}{}

	for gameNumber, playerScores := range halfRoundGameScores {
		for _, playerScore := range playerScores {
			argsStatement = append(argsStatement, "(?, ?, ?, ?, ?)")
			args = append(
				args,
				playerScore.GetPlayerID(),
				handID,
				gameNumber,
				playerScore.GetScore(),
				playerScore.GetRanking(),
			)
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO half_round_games (player_id, hand_id, game_number, score, ranking) VALUES %s",
		strings.Join(argsStatement, ", "),
	)

	if _, err := ope.Exec(c, query, args...); err != nil {
		return err
	}

	return nil
}
