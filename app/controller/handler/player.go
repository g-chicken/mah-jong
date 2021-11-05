package handler

import (
	"context"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/proto/app/services/player/v1"
	"github.com/g-chicken/mah-jong/app/usecase"
)

type playerGRPCHander struct {
	playerUC usecase.PlayerUsecase
}

// NewPlayerServiceServer implements PlayerServerServer proto.
func NewPlayerServiceServer(playerUC usecase.PlayerUsecase) player.PlayerServiceServer {
	return &playerGRPCHander{
		playerUC: playerUC,
	}
}

func (h *playerGRPCHander) CreatePlayer(
	c context.Context,
	req *player.CreatePlayerRequest,
) (*player.CreatePlayerResponse, error) {
	name := req.GetName()
	if name == "" {
		return nil, domain.NewInvalidArgumentError("no name")
	}

	p, err := h.playerUC.CreatePlayer(c, name)
	if err != nil {
		return nil, err
	}

	return &player.CreatePlayerResponse{
		Player: &player.Player{
			Id:   uint32(p.GetID()),
			Name: p.GetName(),
		},
	}, nil
}

func (h *playerGRPCHander) FetchPlayers(
	c context.Context,
	req *player.FetchPlayersRequest,
) (*player.FetchPlayersResponse, error) {
	return &player.FetchPlayersResponse{}, nil
}
