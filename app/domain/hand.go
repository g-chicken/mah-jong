package domain

import (
	"context"
	"time"
)

// CreateHand creates a hand.
func CreateHand(c context.Context, timestamp time.Time) (*Hand, error) {
	return repos.handRepo.CreateHand(c, timestamp)
}

// GetHands get hands.
func GetHands(c context.Context) ([]*Hand, error) {
	return repos.handRepo.GetHands(c)
}

// GetHandByID get a hand whose ID is handID.
func GetHandByID(c context.Context, handID uint64) (*Hand, error) {
	return repos.handRepo.GetHandByID(c, handID)
}

// Hand is the hand model.
type Hand struct {
	id        uint64
	timestamp time.Time
}

// NewHand create a Hand.
func NewHand(id uint64, timestamp time.Time) *Hand {
	return &Hand{
		id:        id,
		timestamp: timestamp,
	}
}

// GetID return the hand's ID.
func (h *Hand) GetID() uint64 {
	if h == nil {
		return 0
	}

	return h.id
}

// GetTimestamp returns the timestamp.
func (h *Hand) GetTimestamp() time.Time {
	if h == nil {
		return time.Time{}
	}

	return h.timestamp
}

// GetHalfScores get scores of half round game in hand.
func (h *Hand) GetHalfScore(c context.Context) (*HandScore, error) {
	if h == nil {
		return nil, errNilHand
	}

	halfRoundGameScores, err := repos.halfRoundGameRepo.GetHalfRoundGameScoresByHandID(c, h.GetID())
	if err != nil {
		return nil, err
	}

	return NewHandScore(h.GetID(), h.GetTimestamp(), halfRoundGameScores), nil
}

// GetHalfRoundGameScore get a players' score of a half round game.
func (h *Hand) GetHalfRoundGameScore(c context.Context, gameNumber uint32) (HalfRoundGameScore, error) {
	if h == nil {
		return nil, errNilHand
	}

	return repos.halfRoundGameRepo.GetHalfRoundGameScoreByHandIDAndGameNumber(c, h.GetID(), gameNumber)
}

// CreateHalfRoundGameScores creates player scores of the half round game.
func (h *Hand) CreateHalfRoundGameScores(c context.Context, halfRoundGameScores HalfRoundGameScores) error {
	if h == nil {
		return errNilHand
	}

	return repos.halfRoundGameRepo.CreateHalfRoundGames(c, h.GetID(), halfRoundGameScores)
}

// GetParticipatePlayerIDs gets IDs of players who participate the hand.
func (h *Hand) GetParticipatePlayerIDs(c context.Context) ([]uint64, error) {
	if h == nil {
		return nil, errNilHand
	}

	return repos.playerHandRepo.ParticipatePlayersInHand(c, h.GetID())
}
