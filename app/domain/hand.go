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
