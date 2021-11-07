package handler_test

import (
	"os"
	"testing"

	"github.com/g-chicken/mah-jong/app/proto/app/services/hand/v1"
	"github.com/g-chicken/mah-jong/app/proto/app/services/player/v1"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ignoreUnexported = cmpopts.IgnoreUnexported(
		player.CreatePlayerResponse{},
		player.FetchPlayersResponse{},
		player.Player{},
		hand.CreateHandResponse{},
		hand.FetchHandScoreResponse{},
		hand.FetchHandsResponse{},
		hand.Hand{},
		timestamppb.Timestamp{},
	)
	noErrFunc = func(err error) bool { return err != nil }
	errFunc   = func(err error) bool { return err == nil }
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
