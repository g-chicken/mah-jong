package usecase_test

import (
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/usecase"
	"github.com/google/go-cmp/cmp"
)

func TestCreateHandArguments_ToHalfRoundGameScores(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		args      *usecase.CreateHandArguments
		want      domain.HalfRoundGameScores
		playerIDs []uint64
	}{
		{
			name: "normal",
			args: &usecase.CreateHandArguments{
				Timestamp: time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				PlayerScores: []usecase.CreateHandArgumentsPlayerScore{
					{PlayerID: 1, Score: 10, GameNumber: 1},
					{PlayerID: 1, Score: 0, GameNumber: 2},
					{PlayerID: 2, Score: 20, GameNumber: 1},
					{PlayerID: 3, Score: -30, GameNumber: 1},
					{PlayerID: 3, Score: 0, GameNumber: 3},
					{PlayerID: 3, Score: 40, GameNumber: 2},
					{PlayerID: 2, Score: -40, GameNumber: 2},
					{PlayerID: 2, Score: 10, GameNumber: 3},
					{PlayerID: 1, Score: -10, GameNumber: 3},
				},
			},
			want: domain.HalfRoundGameScores{
				1: []*domain.PlayerScore{
					domain.NewPlayerScore(2, 20, 1),
					domain.NewPlayerScore(1, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: []*domain.PlayerScore{
					domain.NewPlayerScore(3, 40, 1),
					domain.NewPlayerScore(1, 0, 2),
					domain.NewPlayerScore(2, -40, 3),
				},
				3: []*domain.PlayerScore{
					domain.NewPlayerScore(2, 10, 1),
					domain.NewPlayerScore(3, 0, 2),
					domain.NewPlayerScore(1, -10, 3),
				},
			},
			playerIDs: []uint64{2, 1, 3},
		},
		{
			name: "same player ID",
			args: &usecase.CreateHandArguments{
				Timestamp: time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				PlayerScores: []usecase.CreateHandArgumentsPlayerScore{
					{PlayerID: 1, Score: 0, GameNumber: 1},
					{PlayerID: 1, Score: 100, GameNumber: 3},
					{PlayerID: 1, Score: 5, GameNumber: 3},
					{PlayerID: 1, Score: -20, GameNumber: 3},
					{PlayerID: 1, Score: 40, GameNumber: 3},
					{PlayerID: 1, Score: -80, GameNumber: 1},
					{PlayerID: 1, Score: -110, GameNumber: 3},
				},
			},
			want: domain.HalfRoundGameScores{
				1: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 0, 1),
					domain.NewPlayerScore(1, -80, 2),
				},
				3: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 100, 1),
					domain.NewPlayerScore(1, 40, 2),
					domain.NewPlayerScore(1, 5, 3),
					domain.NewPlayerScore(1, -20, 4),
					domain.NewPlayerScore(1, -110, 5),
				},
			},
			playerIDs: []uint64{1},
		},
		{
			name: "empty",
			args: &usecase.CreateHandArguments{
				Timestamp:    time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				PlayerScores: []usecase.CreateHandArgumentsPlayerScore{},
			},
			want:      domain.HalfRoundGameScores{},
			playerIDs: []uint64{},
		},
		{
			name:      "nil",
			args:      nil,
			want:      domain.HalfRoundGameScores{},
			playerIDs: []uint64{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, playerIDs := tc.args.ToHalfRoundGameScores()

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result(-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.playerIDs, playerIDs, uintArraySort); diff != "" {
				t.Fatalf("unexpected result(-want +got):\n%s", diff)
			}
		})
	}
}
