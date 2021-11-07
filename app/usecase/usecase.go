//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package usecase

import (
	"context"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
)

// ConfigUsecase defines usecase of config.
type ConfigUsecase interface {
	GetConfig(c context.Context) (*domain.Config, error)
}

// PlayerUsecase defines usecase of player.
type PlayerUsecase interface {
	CreatePlayer(c context.Context, name string) (*domain.Player, error)
	FetchPlayers(c context.Context) ([]*domain.Player, error)
}

// CreateHandArguments is a argument of CreateHand method.
type CreateHandArguments struct {
	Timestamp    time.Time
	PlayerScores []CreateHandArgumentsPlayerScore
}

func (args *CreateHandArguments) validate() bool {
	return args != nil && len(args.PlayerScores) != 0 && !args.Timestamp.IsZero()
}

func (args *CreateHandArguments) ToHalfRoundGameScores() (domain.HalfRoundGameScores, []uint64) {
	if args == nil {
		return domain.HalfRoundGameScores{}, []uint64{}
	}

	indexesOrderedScore := args.getOrderedIndexes()

	return args.toHalfRoundGameScores(indexesOrderedScore)
}

// getOrderedIndexes's return value is [game number] = PlayerScores indexes.
func (args *CreateHandArguments) getOrderedIndexes() map[uint32][]int {
	indexesOrderedScore := map[uint32][]int{} // [game number] = PlayerScores indexes

	for i, playerScore := range args.PlayerScores {
		gameNumber := playerScore.GameNumber

		playerScoresIndexes, ok := indexesOrderedScore[gameNumber]
		if !ok {
			indexesOrderedScore[gameNumber] = []int{i}

			continue
		}

		playerScoresIndexes = append(playerScoresIndexes, i)

		for j := len(playerScoresIndexes) - 1; j > 0; j-- {
			previousPlayerScore := args.PlayerScores[playerScoresIndexes[j-1]]

			if previousPlayerScore.Score < playerScore.Score {
				playerScoresIndexes[j], playerScoresIndexes[j-1] = playerScoresIndexes[j-1], playerScoresIndexes[j]
			} else {
				break
			}
		}

		indexesOrderedScore[gameNumber] = playerScoresIndexes
	}

	return indexesOrderedScore
}

func (args *CreateHandArguments) toHalfRoundGameScores(
	indexesOrderedScore map[uint32][]int,
) (domain.HalfRoundGameScores, []uint64) {
	halfRoundGameScores := domain.HalfRoundGameScores{}
	playerIDs := []uint64{}

	for gameNumber, playerScoresIndexes := range indexesOrderedScore {
		domainPlayerScores := make([]*domain.PlayerScore, 0, len(playerScoresIndexes))

		for ranking, index := range playerScoresIndexes {
			playerScore := args.PlayerScores[index]
			p := domain.NewPlayerScore(playerScore.PlayerID, playerScore.Score, uint32(ranking+1))

			domainPlayerScores = append(domainPlayerScores, p)

			notExist := true

			for _, playerID := range playerIDs {
				if playerScore.PlayerID == playerID {
					notExist = false

					break
				}
			}

			if notExist {
				playerIDs = append(playerIDs, playerScore.PlayerID)
			}
		}

		halfRoundGameScores[gameNumber] = domainPlayerScores
	}

	return halfRoundGameScores, playerIDs
}

// CreateHandArgumentsPlayerScore is a player score used CreateHandArguments.
type CreateHandArgumentsPlayerScore struct {
	PlayerID   uint64
	Score      int
	GameNumber uint32
}

// HandUsecase defines usecase of hand.
type HandUsecase interface {
	CreateHand(c context.Context, args *CreateHandArguments) (*domain.Hand, []uint64, error)
	FetchHands(c context.Context) ([]*domain.Hand, map[uint64][]uint64 /*  [hand ID] = {plyer IDs} */, error)
}
