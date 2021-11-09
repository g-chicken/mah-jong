package rdb_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/google/go-cmp/cmp"
)

func TestHalfRoundGameRepository_CreateHalfRoundGames(t *testing.T) {
	type want struct {
		PlayerID   uint64
		HandID     uint64
		GameNumber uint32
		Score      int
		Ranking    uint32
	}

	testCases := []struct {
		name                string
		handID              uint64
		halfRoundGameScores domain.HalfRoundGameScores
		want                []want
		errFunc             func(error) bool
	}{
		{
			name:   "success",
			handID: 3,
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(2, 10, 3),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 30, 1),
					domain.NewPlayerScore(5, -60, 4),
				},
				2: {
					domain.NewPlayerScore(2, 30, 1),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 10, 3),
					domain.NewPlayerScore(5, -60, 4),
				},
			},
			want: []want{
				{PlayerID: 2, HandID: 3, GameNumber: 1, Score: 10, Ranking: 3},
				{PlayerID: 3, HandID: 3, GameNumber: 1, Score: 20, Ranking: 2},
				{PlayerID: 4, HandID: 3, GameNumber: 1, Score: 30, Ranking: 1},
				{PlayerID: 5, HandID: 3, GameNumber: 1, Score: -60, Ranking: 4},
				{PlayerID: 2, HandID: 3, GameNumber: 2, Score: 30, Ranking: 1},
				{PlayerID: 3, HandID: 3, GameNumber: 2, Score: 20, Ranking: 2},
				{PlayerID: 4, HandID: 3, GameNumber: 2, Score: 10, Ranking: 3},
				{PlayerID: 5, HandID: 3, GameNumber: 2, Score: -60, Ranking: 4},
			},
			errFunc: notErrFunc,
		},
		{
			name:    "empty",
			handID:  3,
			errFunc: notErrFunc,
		},
		{
			name:   "invalid halfRoundGameScores",
			handID: 3,
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(2, 10, 3),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(3, 30, 1),
					domain.NewPlayerScore(5, -60, 4),
				},
				2: {
					domain.NewPlayerScore(2, 30, 1),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 10, 3),
					domain.NewPlayerScore(5, -60, 4),
				},
			},
			errFunc: func(err error) bool { return !errors.As(err, &domain.InvalidArgumentError{}) },
		},
		{
			name:   "not found hand ID",
			handID: 99,
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(2, 10, 3),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 30, 1),
					domain.NewPlayerScore(5, -60, 4),
				},
				2: {
					domain.NewPlayerScore(2, 30, 1),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 10, 3),
					domain.NewPlayerScore(5, -60, 4),
				},
			},
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.IllegalForeignKeyConstraintError{})
			},
		},
		{
			name:   "not player ID",
			handID: 3,
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(99, 10, 3),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 30, 1),
					domain.NewPlayerScore(5, -60, 4),
				},
				2: {
					domain.NewPlayerScore(2, 30, 1),
					domain.NewPlayerScore(3, 20, 2),
					domain.NewPlayerScore(4, 10, 3),
					domain.NewPlayerScore(5, -60, 4),
				},
			},
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.IllegalForeignKeyConstraintError{})
			},
		},
		{
			name:   "conflict",
			handID: 1,
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(1, 0, 1),
				},
			},
			errFunc: func(err error) bool { return !errors.As(err, &domain.ConflictError{}) },
		},
	}

	for _, tc := range testCases {
		tc := tc

		existTest := func(t *testing.T, c context.Context) {
			t.Helper()

			query := "SELECT player_id, hand_id, game_number, score, ranking" +
				" FROM half_round_games WHERE hand_id = ? ORDER BY game_number"
			args := []interface{}{tc.handID}
			ope := rdbDetectorRepo.GetRDBOperator(c)
			got := make([]want, 0, len(tc.want))
			scanFunc := func(rows *sql.Rows) error {
				var (
					playerID   uint64
					handID     uint64
					gameNumber uint32
					score      int
					ranking    uint32
				)

				for rows.Next() {
					_ = rows.Scan(&playerID, &handID, &gameNumber, &score, &ranking)

					got = append(
						got,
						want{
							PlayerID:   playerID,
							HandID:     handID,
							GameNumber: gameNumber,
							Score:      score,
							Ranking:    ranking,
						},
					)
				}

				return nil
			}

			_ = ope.Select(c, query, args, scanFunc)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		}

		t.Run(tc.name, func(t *testing.T) {
			defer initializeHalfRoundGames()

			c := context.Background()

			testHalfRoundGameRepositoryCreateHalfRoundGamesNomal(
				c, t, tc.handID, tc.halfRoundGameScores, tc.errFunc,
			)

			if len(tc.want) > 0 {
				existTest(t, c)
			}
		})

		t.Run(tc.name+"(transaction)", func(t *testing.T) {
			defer initializeHalfRoundGames()

			c := context.Background()

			testHalfRoundGameRepositoryCreateHalfRoundGamesTransaction(
				c, t, tc.handID, tc.halfRoundGameScores, tc.errFunc,
			)

			if len(tc.want) > 0 {
				existTest(t, c)
			}
		})
	}
}

func testHalfRoundGameRepositoryCreateHalfRoundGamesNomal(
	c context.Context,
	t *testing.T,
	handID uint64,
	halfRoundGameScores domain.HalfRoundGameScores,
	errFunc func(error) bool,
) {
	t.Helper()

	repo := rdb.NewHalfRoundGameRepository(rdbDetectorRepo)
	err := repo.CreateHalfRoundGames(c, handID, halfRoundGameScores)

	if errFunc(err) {
		t.Fatalf("unexpected error (error = %v)", err)
	}
}

func testHalfRoundGameRepositoryCreateHalfRoundGamesTransaction(
	c context.Context,
	t *testing.T,
	handID uint64,
	halfRoundGameScores domain.HalfRoundGameScores,
	errFunc func(error) bool,
) {
	t.Helper()

	if err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
		repo := rdb.NewHalfRoundGameRepository(rdbDetectorRepo)
		err := repo.CreateHalfRoundGames(c, handID, halfRoundGameScores)

		if errFunc(err) {
			return fmt.Errorf("unexpected error (error = %w)", err)
		}

		return nil
	}); err != nil {
		t.Fatalf("should not be error but %v", err)
	}
}

func TestHalfRoundGameRepository_GetHalfRoundGameScoresByHandID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		handID  uint64
		want    domain.HalfRoundGameScores
		errFunc func(error) bool
	}{
		{
			name:   "success",
			handID: 1,
			want: domain.HalfRoundGameScores{
				1: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 10, 2),
					domain.NewPlayerScore(2, -20, 3),
					domain.NewPlayerScore(3, -30, 4),
					domain.NewPlayerScore(4, 40, 1),
				},
				2: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 14, 2),
					domain.NewPlayerScore(2, -61, 4),
					domain.NewPlayerScore(3, 73, 1),
					domain.NewPlayerScore(4, -26, 3),
				},
			},
			errFunc: notErrFunc,
		},
		{
			name:   "three players",
			handID: 2,
			want: domain.HalfRoundGameScores{
				1: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 0, 2),
					domain.NewPlayerScore(2, -31, 3),
					domain.NewPlayerScore(4, 31, 1),
				},
				2: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 25, 1),
					domain.NewPlayerScore(2, -4, 2),
					domain.NewPlayerScore(4, -21, 3),
				},
				3: []*domain.PlayerScore{
					domain.NewPlayerScore(1, 43, 1),
					domain.NewPlayerScore(2, -55, 3),
					domain.NewPlayerScore(4, 12, 2),
				},
			},
			errFunc: notErrFunc,
		},
		{
			name:    "no hand ID",
			handID:  99,
			errFunc: func(err error) bool { return !errors.As(err, &domain.NotFoundError{}) },
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repo := rdb.NewHalfRoundGameRepository(rdbDetectorRepo)
			got, err := repo.GetHalfRoundGameScoresByHandID(context.Background(), tc.handID)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported, uint64KeySort, playerScoreSliceSort); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})

		t.Run(tc.name+"(transaction)", func(t *testing.T) {
			t.Parallel()

			if err := rdbStatementSetRepo.Transaction(context.Background(), func(c context.Context) error {
				repo := rdb.NewHalfRoundGameRepository(rdbDetectorRepo)
				got, err := repo.GetHalfRoundGameScoresByHandID(context.Background(), tc.handID)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported, uint64KeySort, playerScoreSliceSort); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			}); err != nil {
				t.Fatalf("unexpected error (error = %v)", err)
			}
		})
	}
}
