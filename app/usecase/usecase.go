//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package usecase

import (
	"context"

	"github.com/g-chicken/mah-jong/app/domain"
)

// ConfigUsecase defines usecase of config.
type ConfigUsecase interface {
	GetConfig(c context.Context) (*domain.Config, error)
}
