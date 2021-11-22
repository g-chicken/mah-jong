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

// UpdateScoreAndRanking updates score and ranking in hand.
// scores of arguments is [player ID] = score.
func (h *HandScore) UpdateScoreAndRanking(c context.Context, gameNumber uint32, scores map[uint64]int) error {
	if h == nil {
		return errNilHandScore
	}

	halfRoundScore, ok := h.GetHalfGameScores()[gameNumber]
	if !ok {
		return nil
	}

	return halfRoundScore.updateScoresAndRankings(c, h.GetID(), gameNumber, scores)
}
