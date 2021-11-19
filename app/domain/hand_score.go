package domain

import (
	"context"
	"time"
)

// HandScore is a score of hand.
type HandScore struct {
	id                  uint64
	timestamp           time.Time
	halfRoundGameScores HalfRoundGameScores
}

// NewHandScore is a hand score model.
func NewHandScore(
	id uint64, timestamp time.Time, halfRoundGameScores HalfRoundGameScores,
) *HandScore {
	return &HandScore{
		id:                  id,
		timestamp:           timestamp,
		halfRoundGameScores: halfRoundGameScores,
	}
}

// GetID returns the hand ID.
func (h *HandScore) GetID() uint64 {
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

// CreateHalfRoundGameScores creates player scores of the half round game.
func CreateHalfRoundGameScores(
	c context.Context,
	handID uint64,
	halfRoundGameScores HalfRoundGameScores,
) error {
	return repos.halfRoundGameRepo.CreateHalfRoundGames(c, handID, halfRoundGameScores)
}
