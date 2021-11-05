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
