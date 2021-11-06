package rdb

import (
	"context"
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
