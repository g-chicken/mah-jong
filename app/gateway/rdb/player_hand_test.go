package rdb_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/google/go-cmp/cmp"
)

func TestPlayerHandRepository_CreatePlayerHand(t *testing.T) {
	type want struct {
		ID       uint64
		PlayerID uint64
		HandID   uint64
	}

	testCases := []struct {
		name    string
		args    []*domain.CreatePlayerHandArgs
		want    []want
		errFunc func(error) bool
	}{
		{
			name: "success",
			args: []*domain.CreatePlayerHandArgs{
				{PlayerID: 2, HandID: 3},
				{PlayerID: 3, HandID: 3},
				{PlayerID: 4, HandID: 3},
				{PlayerID: 5, HandID: 3},
			},
			want: []want{
				{ID: 8, PlayerID: 2, HandID: 3},
				{ID: 9, PlayerID: 3, HandID: 3},
				{ID: 10, PlayerID: 4, HandID: 3},
				{ID: 11, PlayerID: 5, HandID: 3},
			},
			errFunc: notErrFunc,
		},
		{
			name:    "empty args",
			args:    []*domain.CreatePlayerHandArgs{},
			want:    []want{},
			errFunc: notErrFunc,
		},
		{
			name: "conflict",
			args: []*domain.CreatePlayerHandArgs{
				{PlayerID: 1, HandID: 1},
			},
			errFunc: func(err error) bool { return !errors.As(err, &domain.ConflictError{}) },
		},
		{
			name: "not found player ID",
			args: []*domain.CreatePlayerHandArgs{
				{PlayerID: 99, HandID: 1},
			},
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.IllegalForeignKeyConstraintError{})
			},
		},
		{
			name: "not found hand ID",
			args: []*domain.CreatePlayerHandArgs{
				{PlayerID: 1, HandID: 99},
			},
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.IllegalForeignKeyConstraintError{})
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			defer initializePlayersHands()

			c := context.Background()

			repo := rdb.NewPlayerHandRepository(rdbDetectorRepo)
			err := repo.CreatePlayerHandPaires(c, tc.args)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if err == nil {
				query := "SELECT id, player_id, hand_id FROM players_hands WHERE id IN (?, ?, ?, ?)"
				args := []interface{}{8, 9, 10, 11}
				ope := rdbDetectorRepo.GetRDBOperator(c)

				got := make([]want, 0, 4)

				scanFunc := func(rows *sql.Rows) error {
					var (
						id       uint64
						playerID uint64
						handID   uint64
					)

					for rows.Next() {
						_ = rows.Scan(&id, &playerID, &handID)
						got = append(got, want{ID: id, PlayerID: playerID, HandID: handID})
					}

					return nil
				}

				_ = ope.Select(c, query, args, scanFunc)

				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Fatalf("unexpected result (-want +got):\n%s", diff)
				}
			}
		})
	}
}
