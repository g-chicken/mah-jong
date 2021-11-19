package handler

import (
	"context"

	"github.com/g-chicken/mah-jong/app/proto/app/services/hand/v1"
	"github.com/g-chicken/mah-jong/app/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handGRPCHandler struct {
	handUC usecase.HandUsecase
}

// NewHandServiceServer implements HandServiceServer proto.
func NewHandServiceServer(handUC usecase.HandUsecase) hand.HandServiceServer {
	return &handGRPCHandler{
		handUC: handUC,
	}
}

func (h *handGRPCHandler) CreateHand(
	c context.Context,
	req *hand.CreateHandRequest,
) (*hand.CreateHandResponse, error) {
	playerScores := make([]usecase.CreateHandArgumentsPlayerScore, 0, len(req.GetPlayerScores()))

	for _, ps := range req.GetPlayerScores() {
		playerScores = append(
			playerScores,
			usecase.CreateHandArgumentsPlayerScore{
				PlayerID:   ps.GetPlayerId(),
				Score:      int(ps.GetScore()),
				GameNumber: ps.GetGameNumber(),
			},
		)
	}

	args := &usecase.CreateHandArguments{
		Timestamp:    req.GetTimestamp().AsTime(),
		PlayerScores: playerScores,
	}

	domainHand, playerIDs, err := h.handUC.CreateHand(c, args)
	if err != nil {
		return nil, err
	}

	return &hand.CreateHandResponse{
		Hand: &hand.Hand{
			Id:                   domainHand.GetID(),
			ParticipatePlayerIds: playerIDs,
			Timestamp:            timestamppb.New(domainHand.GetTimestamp()),
		},
	}, nil
}

func (h *handGRPCHandler) FetchHandScore(
	c context.Context,
	req *hand.FetchHandScoreRequest,
) (*hand.FetchHandScoreResponse, error) {
	domainHand, playerIDs, scores, err := h.handUC.FetchHandScore(c, req.GetHandId())
	if err != nil {
		return nil, err
	}

	halfGameScores := map[uint32]*hand.HandScore_HalfGameScore{}

	for gameNumber, playerScores := range scores {
		playerScoresPB := make([]*hand.HandScore_HalfGameScore_PlayerScore, 0, len(playerScores))

		for _, playerScore := range playerScores {
			playerScoresPB = append(
				playerScoresPB,
				&hand.HandScore_HalfGameScore_PlayerScore{
					PlayerId: playerScore.GetPlayerID(),
					Score:    int32(playerScore.GetScore()),
					Ranking:  playerScore.GetRanking(),
				},
			)
		}

		halfGameScores[gameNumber] = &hand.HandScore_HalfGameScore{PlayerScores: playerScoresPB}
	}

	return &hand.FetchHandScoreResponse{
		HandScore: &hand.HandScore{
			Id:                   domainHand.GetID(),
			ParticipatePlayerIds: playerIDs,
			Timestamp:            timestamppb.New(domainHand.GetTimestamp()),
			HalfGameScores:       halfGameScores,
		},
	}, nil
}

func (h *handGRPCHandler) FetchHands(
	c context.Context,
	req *hand.FetchHandsRequest,
) (*hand.FetchHandsResponse, error) {
	hands, playerIDsInHand, err := h.handUC.FetchHands(c)
	if err != nil {
		return nil, err
	}

	handPBs := make([]*hand.Hand, 0, len(hands))

	for _, h := range hands {
		handPB := &hand.Hand{
			Id:                   h.GetID(),
			ParticipatePlayerIds: playerIDsInHand[h.GetID()],
			Timestamp:            timestamppb.New(h.GetTimestamp()),
		}

		handPBs = append(handPBs, handPB)
	}

	return &hand.FetchHandsResponse{Hands: handPBs}, nil
}

func (h *handGRPCHandler) UpdateHandScores(
	c context.Context,
	req *hand.UpdateHandScoresRequest,
) (*hand.UpdateHandScoresResponse, error) {
	return &hand.UpdateHandScoresResponse{}, nil
}
