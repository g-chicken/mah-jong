package rdb_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/g-chicken/mah-jong/app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
)

func TestMain(m *testing.M) {
	before()

	code := m.Run()

	after()

	os.Exit(code)
}

var (
	notErrFunc      = func(err error) bool { return err != nil }
	allowUnexported = cmp.AllowUnexported(domain.Player{})
	rdbGetterRepo   domain.RDBGetterRepository

	closeFunc func()

	playerDTO = []struct {
		ID   uint64
		Name string
	}{
		{
			ID:   1,
			Name: "test",
		},
		{
			ID:   2,
			Name: "hoge",
		},
		{
			ID:   3,
			Name: "foo",
		},
		{
			ID:   4,
			Name: "bar",
		},
		{
			ID:   5,
			Name: "fuga",
		},
	}
)

func before() {
	_ = logger.SetLogger()

	config := domain.NewConfig(8080, "localhost:3306", "mah_jong_test", "app", "hoge", 5*time.Second)
	rdbGetterRepo, closeFunc, _ = rdb.NewRDBGetterRepository(config)

	setPlayers()
}

func after() {
	removeAllPlayers()

	logger.CloseLogger()
	closeFunc()
}

func setPlayers() {
	query := "INSERT INTO players (id, name) VALUES"
	args := []interface{}{}

	for i, dto := range playerDTO {
		args = append(args, dto.ID, dto.Name)
		query += " (?, ?)"

		if i < len(playerDTO)-1 {
			query += ","
		}
	}

	c := context.Background()
	ope := rdbGetterRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}

func removeAllPlayers() {
	c := context.Background()
	ope := rdbGetterRepo.GetRDBOperator(c)

	deleteQuery := "DELETE FROM players"
	_, _ = ope.Exec(c, deleteQuery)

	resetAutoIncrementQuery := "ALTER TABLE players auto_increment = 1"
	_, _ = ope.Exec(c, resetAutoIncrementQuery)
}

func initializePlayers() {
	query := "DELETE FROM players WHERE id NOT IN ("
	args := []interface{}{}

	for i, dto := range playerDTO {
		args = append(args, dto.ID)
		query += "?"

		if i < len(playerDTO)-1 {
			query += ", "
		}
	}

	query += ")"

	c := context.Background()
	ope := rdbGetterRepo.GetRDBOperator(c)
	_, _ = ope.Exec(c, query, args...)
}
