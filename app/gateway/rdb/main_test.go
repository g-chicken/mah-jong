package rdb_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/g-chicken/mah-jong/app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMain(m *testing.M) {
	before()

	code := m.Run()

	after()

	os.Exit(code)
}

var (
	notErrFunc      = func(err error) bool { return err != nil }
	allowUnexported = cmp.AllowUnexported(
		domain.Player{},
		domain.PlayerScore{},
		domain.Hand{},
		domain.HandScore{},
	)
	uint64KeySort        = cmpopts.SortMaps(func(x, y uint64) bool { return x < y })
	playerScoreSliceSort = cmpopts.SortSlices(
		func(x, y *domain.PlayerScore) bool { return x.GetPlayerID() < y.GetPlayerID() },
	)
	rdbStatementSetRepo domain.RDBStatementSetRepository
	rdbDetectorRepo     domain.RDBDetectorRepository

	closeFunc func()

	playerDTO = []struct {
		id   uint64
		name string
	}{
		{
			id:   1,
			name: "test",
		},
		{
			id:   2,
			name: "hoge",
		},
		{
			id:   3,
			name: "foo",
		},
		{
			id:   4,
			name: "bar",
		},
		{
			id:   5,
			name: "fuga",
		},
	}
	handDTO = []struct {
		id       uint64
		gameDate time.Time
	}{
		{
			id:       1,
			gameDate: time.Date(2021, time.November, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			id:       2,
			gameDate: time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
		},
		{
			id:       3,
			gameDate: time.Date(2021, time.November, 8, 0, 0, 0, 0, time.UTC),
		},
	}
	playerHandDTO = []struct {
		id       uint64
		playerID uint64
		handID   uint64
	}{
		{
			id:       1,
			playerID: 1,
			handID:   1,
		},
		{
			id:       2,
			playerID: 2,
			handID:   1,
		},
		{
			id:       3,
			playerID: 3,
			handID:   1,
		},
		{
			id:       4,
			playerID: 4,
			handID:   1,
		},
		{
			id:       5,
			playerID: 1,
			handID:   2,
		},
		{
			id:       6,
			playerID: 2,
			handID:   2,
		},
		{
			id:       7,
			playerID: 4,
			handID:   2,
		},
	}
	halfRoundGameDTO = []struct {
		id         uint64
		playerID   uint64
		handID     uint64
		gameNumber uint32
		score      int
		ranking    uint32
	}{
		{
			id:         1,
			playerID:   1,
			handID:     1,
			gameNumber: 1,
			score:      10,
			ranking:    2,
		},
		{
			id:         2,
			playerID:   2,
			handID:     1,
			gameNumber: 1,
			score:      -20,
			ranking:    3,
		},
		{
			id:         3,
			playerID:   3,
			handID:     1,
			gameNumber: 1,
			score:      -30,
			ranking:    4,
		},
		{
			id:         4,
			playerID:   4,
			handID:     1,
			gameNumber: 1,
			score:      40,
			ranking:    1,
		},
		{
			id:         5,
			playerID:   1,
			handID:     1,
			gameNumber: 2,
			score:      14,
			ranking:    2,
		},
		{
			id:         6,
			playerID:   2,
			handID:     1,
			gameNumber: 2,
			score:      -61,
			ranking:    4,
		},
		{
			id:         7,
			playerID:   3,
			handID:     1,
			gameNumber: 2,
			score:      73,
			ranking:    1,
		},
		{
			id:         8,
			playerID:   4,
			handID:     1,
			gameNumber: 2,
			score:      -26,
			ranking:    3,
		},
		{
			id:         9,
			playerID:   1,
			handID:     2,
			gameNumber: 1,
			score:      0,
			ranking:    2,
		},
		{
			id:         10,
			playerID:   2,
			handID:     2,
			gameNumber: 1,
			score:      -31,
			ranking:    3,
		},
		{
			id:         11,
			playerID:   4,
			handID:     2,
			gameNumber: 1,
			score:      31,
			ranking:    1,
		},
		{
			id:         12,
			playerID:   1,
			handID:     2,
			gameNumber: 2,
			score:      25,
			ranking:    1,
		},
		{
			id:         13,
			playerID:   2,
			handID:     2,
			gameNumber: 2,
			score:      -4,
			ranking:    2,
		},
		{
			id:         14,
			playerID:   4,
			handID:     2,
			gameNumber: 2,
			score:      -21,
			ranking:    3,
		},
		{
			id:         15,
			playerID:   1,
			handID:     2,
			gameNumber: 3,
			score:      43,
			ranking:    1,
		},
		{
			id:         16,
			playerID:   2,
			handID:     2,
			gameNumber: 3,
			score:      -55,
			ranking:    3,
		},
		{
			id:         17,
			playerID:   4,
			handID:     2,
			gameNumber: 3,
			score:      12,
			ranking:    2,
		},
	}
)

func before() {
	_ = logger.SetLogger()

	config := domain.NewConfig(8080, "localhost:3306", "mah_jong_test", "app", "hoge", 5*time.Second)
	rdbStatementSetRepo, rdbDetectorRepo, closeFunc, _ = rdb.NewRDBDetectorRepository(config)

	setPlayers()
	setHands()
	setPlayersHands()
	setHalfRoundGames()
}

func after() {
	removeAllHalfRoundGames()
	removeAllPlayersHands()
	removeAllPlayers()
	removeAllHands()

	logger.CloseLogger()
	closeFunc()
}

func setPlayers() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range playerDTO {
		argsStatement = append(argsStatement, "(?, ?)")
		args = append(args, dto.id, dto.name)
	}

	query := fmt.Sprintf(
		"INSERT INTO players (id, name) VALUES %s",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func removeAllPlayers() {
	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)

	deleteQuery := "DELETE FROM players"
	_, _ = ope.Exec(c, deleteQuery)

	resetAutoIncrementQuery := "ALTER TABLE players auto_increment = 1"
	_, _ = ope.Exec(c, resetAutoIncrementQuery)
}

func initializePlayers() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range playerDTO {
		argsStatement = append(argsStatement, "?")
		args = append(args, dto.id)
	}

	query := fmt.Sprintf(
		"DELETE FROM players WHERE id NOT IN (%s)",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func setHands() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range handDTO {
		argsStatement = append(argsStatement, "(?, ?)")
		args = append(args, dto.id, dto.gameDate)
	}

	query := fmt.Sprintf(
		"INSERT INTO hands (id, game_date) VALUES %s",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func removeAllHands() {
	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)

	deleteQuery := "DELETE FROM hands"
	_, _ = ope.Exec(c, deleteQuery)

	resetAutoIncrementQuery := "ALTER TABLE hands auto_increment = 1"
	_, _ = ope.Exec(c, resetAutoIncrementQuery)
}

func initializeHands() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range handDTO {
		argsStatement = append(argsStatement, "?")
		args = append(args, dto.id)
	}

	query := fmt.Sprintf(
		"DELETE FROM hands WHERE id NOT IN (%s)",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func setPlayersHands() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range playerHandDTO {
		argsStatement = append(argsStatement, "(?, ?, ?)")
		args = append(args, dto.id, dto.playerID, dto.handID)
	}

	query := fmt.Sprintf(
		"INSERT INTO players_hands (id, player_id, hand_id) VALUES %s",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func removeAllPlayersHands() {
	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)

	deleteQuery := "DELETE FROM players_hands"
	_, _ = ope.Exec(c, deleteQuery)

	resetAutoIncrementQuery := "ALTER TABLE players_hands auto_increment = 1"
	_, _ = ope.Exec(c, resetAutoIncrementQuery)
}

func initializePlayersHands() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range playerHandDTO {
		argsStatement = append(argsStatement, "?")
		args = append(args, dto.id)
	}

	query := fmt.Sprintf(
		"DELETE FROM players_hands WHERE id NOT IN (%s)",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func setHalfRoundGames() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range halfRoundGameDTO {
		argsStatement = append(argsStatement, "(?, ?, ?, ?, ?, ?)")
		args = append(args, dto.id, dto.playerID, dto.handID, dto.gameNumber, dto.ranking, dto.score)
	}

	query := fmt.Sprintf(
		"INSERT INTO half_round_games (id, player_id, hand_id, game_number, ranking, score) VALUES %s",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func removeAllHalfRoundGames() {
	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)

	deleteQuery := "DELETE FROM half_round_games"
	_, _ = ope.Exec(c, deleteQuery)

	resetAutoIncrementQuery := "ALTER TABLE half_round_games auto_increment = 1"
	_, _ = ope.Exec(c, resetAutoIncrementQuery)
}

func initializeHalfRoundGames() {
	argsStatement := []string{}
	args := []interface{}{}

	for _, dto := range halfRoundGameDTO {
		argsStatement = append(argsStatement, "?")
		args = append(args, dto.id)
	}

	query := fmt.Sprintf(
		"DELETE FROM half_round_games WHERE id NOT IN (%s)",
		strings.Join(argsStatement, ", "),
	)

	c := context.Background()
	ope := rdbDetectorRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}
