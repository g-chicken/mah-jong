package rdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
)

type handRepository struct {
	repo domain.RDBDetectorRepository
}

// NewHandRepository implements domain.HandRepository.
func NewHandRepository(repo domain.RDBDetectorRepository) domain.HandRepository {
	return &handRepository{
		repo: repo,
	}
}

func (r *handRepository) CreateHand(c context.Context, timestamp time.Time) (*domain.Hand, error) {
	ope := r.repo.GetRDBOperator(c)

	query := "INSERT INTO hands (game_date) VALUE (?)"
	args := []interface{}{timestamp}

	result, err := ope.Exec(c, query, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return domain.NewHand(uint64(id), timestamp), nil
}

func (r *handRepository) GetHands(c context.Context) ([]*domain.Hand, error) {
	ope := r.repo.GetRDBOperator(c)

	query := "SELECT id, game_date FROM hands"
	results := make([]*domain.Hand, 0)
	scanFunc := func(rows *sql.Rows) error {
		var (
			id        uint64
			timestamp time.Time
		)

		for rows.Next() {
			if err := rows.Scan(&id, &timestamp); err != nil {
				return err
			}

			hand := domain.NewHand(id, timestamp)

			results = append(results, hand)
		}

		return nil
	}

	if err := ope.Select(c, query, []interface{}{}, scanFunc); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *handRepository) GetHandByID(c context.Context, id uint64) (*domain.Hand, error) {
	query := "SELECT id, game_date FROM hands WHERE id = ?"
	args := []interface{}{id}

	var (
		handID    uint64
		timestamp time.Time
	)

	ope := r.repo.GetRDBOperator(c)

	if err := ope.Get(c, query, args, &handID, &timestamp); err != nil {
		return nil, err
	}

	return domain.NewHand(handID, timestamp), nil
}
