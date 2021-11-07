package rdb

import (
	"context"
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
