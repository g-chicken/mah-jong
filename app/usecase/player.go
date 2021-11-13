package usecase

import (
	"context"
	"errors"

	"github.com/g-chicken/mah-jong/app/domain"
)

type playerUC struct{}

// NewPlayerUsecase implements PlayerUsecase.
func NewPlayerUsecase() PlayerUsecase {
	return &playerUC{}
}

func (uc *playerUC) CreatePlayer(c context.Context, name string) (*domain.Player, error) {
	player, err := domain.GetPlayerByName(c, name)

	switch {
	case err == nil:
		return player, nil
	case errors.As(err, &domain.NotFoundError{}):
		return domain.CreatePlayer(c, name)
	default:
		return nil, err
	}
}

func (uc *playerUC) FetchPlayers(c context.Context) ([]*domain.Player, error) {
	return domain.GetPlayers(c)
}

func (uc *playerUC) UpdatePlayer(c context.Context, id uint64, name string) (*domain.Player, error) {
	player, err := domain.GetPlayerByID(c, id)
	if err != nil {
		return nil, err
	}

	if err := player.UpdateName(c, name); err != nil {
		return nil, err
	}

	return player, nil
}

func (uc *playerUC) DeletePlayer(c context.Context, id uint64) error {
	player, err := domain.GetPlayerByID(c, id)
	if err != nil {
		return err
	}

	return player.Delete(c)
}
