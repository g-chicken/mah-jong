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
			return hand.CreateHalfRoundGameScores(c, halfRoundGameScores)
		},
	); err != nil {
		return nil, nil, err
	}

	return hand, playerIDs, nil
}

func (uc *handUC) FetchHandScore(
	c context.Context, handID uint64,
) (*domain.HandScore, []uint64, error) {
	hand, err := domain.GetHandByID(c, handID)
	if err != nil {
		return nil, nil, err
	}

	playerIDs, err := hand.GetParticipatePlayerIDs(c)
	if err != nil {
		return nil, nil, err
	}

	handScore, err := hand.GetHalfScore(c)
	if err != nil {
		return nil, nil, err
	}

	return handScore, playerIDs, nil
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
		playerIDs, err := hand.GetParticipatePlayerIDs(c)
		if err != nil {
			return nil, nil, err
		}

		playerIDsInHand[hand.GetID()] = playerIDs
	}

	return hands, playerIDsInHand, nil
}

func (uc *handUC) UpdateHandScore(
	c context.Context, args *UpdateHandScoreArguments,
) (*domain.HandScore, []uint64, error) {
	if args == nil {
		return nil, nil, domain.NewInvalidArgumentError("arguments is nil")
	}

	hand, err := domain.GetHandByID(c, args.HandID)
	if err != nil {
		return nil, nil, err
	}

	participatePlayerIDs, err := hand.GetParticipatePlayerIDs(c)
	if err != nil {
		return nil, nil, err
	}

	handScore, err := hand.GetHalfScore(c)
	if err != nil {
		return nil, nil, err
	}

	if len(args.PlayerScores) == 0 {
		return handScore, participatePlayerIDs, nil
	}

	if err := domain.Transaction(c, func(c context.Context) error {
		for gameNumber, playerScores := range args.PlayerScores {
			scores := map[uint64]int{}

			for _, playerScore := range playerScores {
				scores[playerScore.PlayerID] = playerScore.Score
			}

			if err := handScore.UpdateScoreAndRanking(c, gameNumber, scores); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

	return handScore, participatePlayerIDs, nil
}
