package rdb

import (
	"context"
	"database/sql"

	"github.com/g-chicken/mah-jong/app/domain"
)

type playerRepository struct {
	repo domain.RDBDetectorRepository
}

// NewPlayerRepository implements domain.PlayerRepository.
func NewPlayerRepository(repo domain.RDBDetectorRepository) domain.PlayerRepository {
	return &playerRepository{
		repo: repo,
	}
}

func (r *playerRepository) CreatePlayer(c context.Context, name string) (*domain.Player, error) {
	ope := r.repo.GetRDBOperator(c)

	query := "INSERT INTO players (name) VALUE (?)"
	args := []interface{}{name}

	result, err := ope.Exec(c, query, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return domain.NewPlayer(uint64(id), name), nil
}

func (r *playerRepository) GetPlayerByID(c context.Context, id uint64) (*domain.Player, error) {
	query := "SELECT id, name FROM players WHERE id = ?"
	args := []interface{}{id}

	return r.getPlayer(c, query, args)
}

func (r *playerRepository) GetPlayerByName(c context.Context, name string) (*domain.Player, error) {
	query := "SELECT id, name FROM players WHERE name = ?"
	args := []interface{}{name}

	return r.getPlayer(c, query, args)
}

func (r *playerRepository) getPlayer(
	c context.Context, query string, args []interface{},
) (*domain.Player, error) {
	ope := r.repo.GetRDBOperator(c)

	var (
		id         uint64
		playerName string
	)

	if err := ope.Get(c, query, args, &id, &playerName); err != nil {
		return nil, err
	}

	return domain.NewPlayer(id, playerName), nil
}

func (r *playerRepository) GetPlayers(c context.Context) ([]*domain.Player, error) {
	ope := r.repo.GetRDBOperator(c)

	query := "SELECT id, name FROM players ORDER BY id"
	args := []interface{}{}
	players := make([]*domain.Player, 0)
	scanFunc := func(rows *sql.Rows) error {
		var (
			id   uint64
			name string
		)

		for rows.Next() {
			if err := rows.Scan(&id, &name); err != nil {
				return err
			}

			players = append(players, domain.NewPlayer(id, name))
		}

		return nil
	}

	if err := ope.Select(c, query, args, scanFunc); err != nil {
		return nil, err
	}

	return players, nil
}
