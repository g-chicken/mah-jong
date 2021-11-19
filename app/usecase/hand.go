package usecase

import (
	"context"

	"github.com/g-chicken/mah-jong/app/domain"
)

type handUC struct{}

// NewHandUsecase implements HandUsecase.
func NewHandUsecase() HandUsecase {
	return &handUC{}
}

func (uc *handUC) CreateHand(
	c context.Context, args *CreateHandArguments,
) (*domain.Hand, []uint64, error) {
	if !args.validate() {
		return nil, nil, domain.NewInvalidArgumentError("invalid argument")
	}

	halfRoundGameScores, playerIDs := args.ToHalfRoundGameScores()
	if !halfRoundGameScores.Validate() {
		return nil, nil, domain.NewInvalidArgumentError("invalid scores")
	}

	for _, playerID := range playerIDs {
		_, err := domain.GetPlayerByID(c, playerID)
		if err != nil {
			return nil, nil, err
		}
	}

	var hand *domain.Hand

	if err := domain.Transaction(
		c,
		func(c context.Context) error {
			var err error

			// create hand
			hand, err = domain.CreateHand(c, args.Timestamp)
			if err != nil {
				return err
			}

			// create players_hands
			pairs := make([]*domain.CreatePlayerHandArgs, 0, len(playerIDs))

			for _, playerID := range playerIDs {
				pairs = append(
					pairs,
					&domain.CreatePlayerHandArgs{PlayerID: playerID, HandID: hand.GetID()},
				)
			}

			if err := domain.CreatePlayerHandPairs(c, pairs); err != nil {
				return err
			}

			// create player scores
			return domain.CreateHalfRoundGameScores(c, hand.GetID(), halfRoundGameScores)
		},
	); err != nil {
		return nil, nil, err
	}

	return hand, playerIDs, nil
}

func (uc *handUC) FetchHandScore(
	c context.Context, handID uint64,
) (*domain.Hand, []uint64, domain.HalfRoundGameScores, error) {
	hand, err := domain.GetHandByID(c, handID)
	if err != nil {
		return nil, nil, nil, err
	}

	playerIDs, err := domain.ParticipatePlayersInHand(c, handID)
	if err != nil {
		return nil, nil, nil, err
	}

	handScore, err := hand.GetHalfScore(c)
	if err != nil {
		return nil, nil, nil, err
	}

	return hand, playerIDs, handScore.GetHalfGameScores(), nil
}

func (uc *handUC) FetchHands(
	c context.Context,
) ([]*domain.Hand, map[uint64][]uint64 /* [hand ID] = {plyer IDs} */, error) {
	hands, err := domain.GetHands(c)
	if err != nil {
		return nil, nil, err
	}

	playerIDsInHand := map[uint64][]uint64{}

	for _, hand := range hands {
		playerIDs, err := domain.ParticipatePlayersInHand(c, hand.GetID())
		if err != nil {
			return nil, nil, err
		}

		playerIDsInHand[hand.GetID()] = playerIDs
	}

	return hands, playerIDsInHand, nil
}
