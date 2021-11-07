//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package domain

import (
	"context"
	"database/sql"
	"time"
)

// Transaction set transaction statements.
func Transaction(c context.Context, transaction func(context.Context) error) error {
	return repos.rdbStatementSetRepo.Transaction(c, transaction)
}

// PlayerRepository defines to operate the players table.
type PlayerRepository interface {
	CreatePlayer(c context.Context, name string) (*Player, error)
	GetPlayerByID(c context.Context, id uint64) (*Player, error)
	GetPlayerByName(c context.Context, name string) (*Player, error)
	GetPlayers(c context.Context) ([]*Player, error)
}

// HalfRoundGameRepository defines to operate the half_round_games table.
type HalfRoundGameRepository interface {
	CreateHalfRoundGames(
		c context.Context, handID uint64, halfRoundGameScores HalfRoundGameScores,
	) error
}

// HandRepository defines to operate the hands table.
type HandRepository interface {
	CreateHand(c context.Context, timestamp time.Time) (*Hand, error)
	GetHands(c context.Context) ([]*Hand, error)
}

// CreatePlayerHandArgs is a argument of CreatePlayerHand method.
type CreatePlayerHandArgs struct {
	PlayerID uint64
	HandID   uint64
}

// PlayerHandRepository defines to operate the players_hands table.
type PlayerHandRepository interface {
	CreatePlayerHandPairs(c context.Context, args []*CreatePlayerHandArgs) error
	ParticipatePlayersInHand(c context.Context, handID uint64) ([]uint64, error)
}

// RDBStatementSetRepository defines to set statement. (for example, transaction.)
type RDBStatementSetRepository interface {
	Transaction(c context.Context, f func(c context.Context) error) error
}

// RDBDetectorRepositry defines to fetch DB structure.
type RDBDetectorRepository interface {
	GetRDBOperator(c context.Context) RDBOperator
}

// RDBOperator defines operation of RDB.
type RDBOperator interface {
	Get(c context.Context, query string, args []interface{}, dist ...interface{}) error
	Select(c context.Context, query string, args []interface{}, scanFunc func(*sql.Rows) error) error
	Exec(c context.Context, query string, args ...interface{}) (sql.Result, error)
}

type repositories struct {
	playerRepo          PlayerRepository
	handRepo            HandRepository
	halfRoundGameRepo   HalfRoundGameRepository
	playerHandRepo      PlayerHandRepository
	rdbStatementSetRepo RDBStatementSetRepository
}

var repos *repositories

// SetRepositories set the repos variable.
func SetRepositories(
	playerRepo PlayerRepository,
	handRepo HandRepository,
	halfRoundGameRepo HalfRoundGameRepository,
	playerHandRepo PlayerHandRepository,
	rdbStatementSetRepo RDBStatementSetRepository,
) {
	repos = &repositories{
		playerRepo:          playerRepo,
		handRepo:            handRepo,
		halfRoundGameRepo:   halfRoundGameRepo,
		playerHandRepo:      playerHandRepo,
		rdbStatementSetRepo: rdbStatementSetRepo,
	}
}

// ConfigRepository defines processes for config.
type ConfigRepository interface {
	GetConfig(c context.Context) (*Config, error)
}
