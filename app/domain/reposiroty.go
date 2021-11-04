//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package domain

import "context"

// PlayerRepository defines processes for players.
type PlayerRepository interface {
	CreatePlayer(c context.Context, name string) (*Player, error)
	GetPlayerByName(c context.Context, name string) (*Player, error)
}

type repositories struct {
	PlayerRepository
}

var repos *repositories

// SetRepositories set the repos variable.
func SetRepositories(
	playerRepo PlayerRepository,
) {
	repos = &repositories{
		PlayerRepository: playerRepo,
	}
}

// ConfigRepository defines processes for config.
type ConfigRepository interface {
	GetConfig(c context.Context) (*Config, error)
}
