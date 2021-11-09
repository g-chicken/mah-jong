package rdb

import (
	"context"
	"database/sql"
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

func (r *halfRoundGameRepository) GetHalfRoundGameScoresByHandID(
	c context.Context, handID uint64,
) (domain.HalfRoundGameScores, error) {
	query := "SELECT player_id, game_number, score, ranking FROM half_round_games WHERE hand_id = ?"
	args := []interface{}{handID}

	halfRoundGameScores := domain.HalfRoundGameScores{}

	scanFunc := func(rows *sql.Rows) error {
		var (
			playerID   uint64
			gameNumber uint32
			score      int
			ranking    uint32
		)

		for rows.Next() {
			if err := rows.Scan(&playerID, &gameNumber, &score, &ranking); err != nil {
				return err
			}

			playerScores, ok := halfRoundGameScores[gameNumber]
			if !ok {
				halfRoundGameScores[gameNumber] = []*domain.PlayerScore{domain.NewPlayerScore(playerID, score, ranking)}

				continue
			}

			playerScores = append(playerScores, domain.NewPlayerScore(playerID, score, ranking))
			halfRoundGameScores[gameNumber] = playerScores
		}

		return nil
	}

	ope := r.repo.GetRDBOperator(c)

	if err := ope.Select(c, query, args, scanFunc); err != nil {
		return nil, err
	}

	if len(halfRoundGameScores) == 0 {
		return nil, domain.NewNotFoundError("no scores of half round game")
	}

	return halfRoundGameScores, nil
}
