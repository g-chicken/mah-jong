//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package domain

import (
	"context"
	"database/sql"
)

// PlayerRepository defines processes for players.
type PlayerRepository interface {
	CreatePlayer(c context.Context, name string) (*Player, error)
	GetPlayerByName(c context.Context, name string) (*Player, error)
	GetPlayers(c context.Context) ([]*Player, error)
}

// RDBGetterRepositry defines to fetch DB structure.
type RDBGetterRepository interface {
	GetRDBOperator(c context.Context) RDBOperator
}

// RDBOperator defines operation of RDB.
type RDBOperator interface {
	Get(c context.Context, query string, args []interface{}, dist ...interface{}) error
	Select(c context.Context, query string, args []interface{}, scanFunc func(*sql.Rows) error) error
	Exec(c context.Context, query string, args ...interface{}) (sql.Result, error)
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
