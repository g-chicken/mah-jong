package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/g-chicken/mah-jong/app/domain"
)

type playerHandRepository struct {
	repo domain.RDBDetectorRepository
}

// NewPlayerHandRepository implements domain.PlayerHandRepository.
func NewPlayerHandRepository(repo domain.RDBDetectorRepository) domain.PlayerHandRepository {
	return &playerHandRepository{
		repo: repo,
	}
}

func (r *playerHandRepository) CreatePlayerHandPairs(
	c context.Context, args []*domain.CreatePlayerHandArgs,
) error {
	if len(args) == 0 {
		return nil
	}

	ope := r.repo.GetRDBOperator(c)

	argsStatemant := []string{}
	dbArgs := make([]interface{}, 0, len(args))

	for _, arg := range args {
		argsStatemant = append(argsStatemant, "(?, ?)")
		dbArgs = append(dbArgs, arg.PlayerID, arg.HandID)
	}

	query := fmt.Sprintf(
		"INSERT INTO players_hands (player_id, hand_id) VALUES %s",
		strings.Join(argsStatemant, ", "),
	)

	if _, err := ope.Exec(c, query, dbArgs...); err != nil {
		return err
	}

	return nil
}

func (r *playerHandRepository) ParticipatePlayersInHand(c context.Context, handID uint64) ([]uint64, error) {
	ope := r.repo.GetRDBOperator(c)

	query := "SELECT DISTINCT player_id FROM players_hands WHERE hand_id = ? ORDER BY player_id"
	args := []interface{}{handID}

	playerIDs := make([]uint64, 0)

	scanFunc := func(rows *sql.Rows) error {
		var id uint64

		for rows.Next() {
			if err := rows.Scan(&id); err != nil {
				return err
			}

			playerIDs = append(playerIDs, id)
		}

		return nil
	}

	if err := ope.Select(c, query, args, scanFunc); err != nil {
		return nil, err
	}

	return playerIDs, nil
}
