package handler_test

import (
	"os"
	"testing"

	"github.com/g-chicken/mah-jong/app/proto/app/services/player/v1"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var (
	ignoreUnexported = cmpopts.IgnoreUnexported(
		player.CreatePlayerResponse{},
		player.FetchPlayersResponse{},
		player.Player{},
	)
	noErrFunc = func(err error) bool { return err != nil }
	errFunc   = func(err error) bool { return err == nil }
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
