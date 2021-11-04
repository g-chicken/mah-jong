package usecase

import (
	"context"

	"github.com/g-chicken/mah-jong/app/domain"
)

type configUC struct {
	repo domain.ConfigRepository
}

// NewConfigUsecase implements ConfigUsecase.
func NewConfigUsecase(repo domain.ConfigRepository) ConfigUsecase {
	return &configUC{repo: repo}
}

func (uc *configUC) GetConfig(c context.Context) (*domain.Config, error) {
	return uc.repo.GetConfig(c)
}
