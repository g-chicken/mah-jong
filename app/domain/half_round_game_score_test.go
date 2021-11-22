package domain_test

import (
	"context"
	"errors"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/google/go-cmp/cmp"
)

func TestHalfRoundGameScore_updateScoresAndRankings(t *testing.T) {
	testCases := []struct {
		name               string
		handID             uint64
		gameNumber         uint32
		scores             map[uint64]int
		halfRoundGameScore domain.HalfRoundGameScore
		setMock            func(*mocks)
		want               domain.HalfRoundGameScore
		err                bool
	}{
		{
			name:       "success",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{1: 15, 77: -38, 23: 13},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {
				c := context.Background()
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(1), 15, uint32(1), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(23), 13, uint32(2), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(58), 10, uint32(3), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(77), -38, uint32(4), uint32(3),
				).Return(nil)
			},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, 15, 1),
				domain.NewPlayerScore(23, 13, 2),
				domain.NewPlayerScore(58, 10, 3),
				domain.NewPlayerScore(77, -38, 4),
			},
		},
		{
			name:       "do not update ranking",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{1: -10, 77: 34, 23: -34},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {
				c := context.Background()
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(1), -10, uint32(3), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(23), -34, uint32(4), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(58), 10, uint32(2), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(77), 34, uint32(1), uint32(3),
				).Return(nil)
			},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -10, 3),
				domain.NewPlayerScore(23, -34, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 34, 1),
			},
		},
		{
			name:       "not updated",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{1: -3, 77: 24, 23: -31},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
		},
		{
			name:       "no player ID",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{99: 100, 0: -20},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
		},
		{
			name:       "scores is empty",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
		},
		{
			name:       "scores is nil",
			handID:     11,
			gameNumber: 3,
			scores:     nil,
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
		},
		{
			name:       "invalid scores",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{1: 15, 77: -39, 23: 13},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			err: true,
		},
		{
			name:               "nil",
			handID:             11,
			gameNumber:         3,
			scores:             map[uint64]int{1: 15, 77: -38, 23: 13},
			halfRoundGameScore: nil,
			setMock:            func(m *mocks) {},
			want:               nil,
			err:                true,
		},
		{
			name:       "error in UpdateScoreAndRanking",
			handID:     11,
			gameNumber: 3,
			scores:     map[uint64]int{1: 15, 77: -38, 23: 13},
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -3, 3),
				domain.NewPlayerScore(23, -31, 4),
				domain.NewPlayerScore(58, 10, 2),
				domain.NewPlayerScore(77, 24, 1),
			},
			setMock: func(m *mocks) {
				c := context.Background()
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(1), 15, uint32(1), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(23), 13, uint32(2), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(58), 10, uint32(3), uint32(3),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(11), uint64(77), -38, uint32(4), uint32(3),
				).Return(errors.New("error"))
			},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, 15, 1),
				domain.NewPlayerScore(23, 13, 2),
				domain.NewPlayerScore(58, 10, 3),
				domain.NewPlayerScore(77, 24, 1),
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			mocks, finish := setRepository(t)
			defer finish()

			tc.setMock(mocks)

			err := tc.halfRoundGameScore.UpdateScoreAndRankings(context.Background(), tc.handID, tc.gameNumber, tc.scores)

			if tc.err && err == nil {
				t.Fatal("should be err but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be err but %v", err)
			}

			if diff := cmp.Diff(tc.halfRoundGameScore, tc.want, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHalfRoundScore_updateRanking(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		halfRoundGameScore domain.HalfRoundGameScore
		want               domain.HalfRoundGameScore
	}{
		{
			name: "basic",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 0),
				domain.NewPlayerScore(2, 50, 0),
				domain.NewPlayerScore(3, 10, 0),
				domain.NewPlayerScore(4, -40, 0),
			},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 3),
				domain.NewPlayerScore(2, 50, 1),
				domain.NewPlayerScore(3, 10, 2),
				domain.NewPlayerScore(4, -40, 4),
			},
		},
		{
			name: "same score",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 0),
				domain.NewPlayerScore(2, 50, 0),
				domain.NewPlayerScore(3, -20, 0),
				domain.NewPlayerScore(4, -40, 0),
			},
			want: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 2),
				domain.NewPlayerScore(2, 50, 1),
				domain.NewPlayerScore(3, -20, 3),
				domain.NewPlayerScore(4, -40, 4),
			},
		},
		{
			name:               "none",
			halfRoundGameScore: domain.HalfRoundGameScore{},
			want:               domain.HalfRoundGameScore{},
		},
		{
			name:               "nil",
			halfRoundGameScore: nil,
			want:               nil,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.halfRoundGameScore.UpdateRanking()

			if diff := cmp.Diff(tc.want, tc.halfRoundGameScore, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHalfRoundGameScore_validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		halfRoundGameScore domain.HalfRoundGameScore
		want               bool
	}{
		{
			name: "success",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 3),
				domain.NewPlayerScore(2, 50, 1),
				domain.NewPlayerScore(3, 10, 2),
				domain.NewPlayerScore(4, -40, 4),
			},
			want: true,
		},
		{
			name: "over 4",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 3),
				domain.NewPlayerScore(2, 50, 1),
				domain.NewPlayerScore(3, 10, 2),
				domain.NewPlayerScore(4, -40, 4),
				domain.NewPlayerScore(5, 0, 5),
			},
			want: false,
		},
		{
			name: "invalid ranking",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, 12, 3),
				domain.NewPlayerScore(2, -36, 4),
				domain.NewPlayerScore(3, -21, 2),
				domain.NewPlayerScore(4, 45, 1),
			},
			want: false,
		},
		{
			name: "sum is not 0",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, -20, 3),
				domain.NewPlayerScore(2, 50, 1),
				domain.NewPlayerScore(3, 11, 2),
				domain.NewPlayerScore(4, -40, 4),
			},
			want: false,
		},
		{
			name: "same player ID",
			halfRoundGameScore: domain.HalfRoundGameScore{
				domain.NewPlayerScore(1, 12, 2),
				domain.NewPlayerScore(2, -36, 4),
				domain.NewPlayerScore(1, -21, 3),
				domain.NewPlayerScore(4, 45, 1),
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.halfRoundGameScore.Validate()
			if tc.want != got {
				t.Fatalf("unexpected result (want : %v, got : %v)", tc.want, got)
			}
		})
	}
}

func TestHalfRoundGameScores_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		halfRoundGameScores domain.HalfRoundGameScores
		want                bool
	}{
		{
			name: "success",
			halfRoundGameScores: map[uint32]domain.HalfRoundGameScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: true,
		},
		{
			name: "over 4",
			halfRoundGameScores: map[uint32]domain.HalfRoundGameScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
					domain.NewPlayerScore(5, 0, 5),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
		{
			name: "invalid ranking",
			halfRoundGameScores: map[uint32]domain.HalfRoundGameScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 3),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 2),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
		{
			name: "sum is not 0",
			halfRoundGameScores: map[uint32]domain.HalfRoundGameScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 11, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
		{
			name: "same player ID",
			halfRoundGameScores: map[uint32]domain.HalfRoundGameScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(1, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.halfRoundGameScores.Validate()
			if tc.want != got {
				t.Fatalf("unexpected result (want : %v, got : %v)", tc.want, got)
			}
		})
	}
}
