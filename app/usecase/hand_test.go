package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestHandUsecase_CreateHand(t *testing.T) {
	testCases := []struct {
		name      string
		args      *usecase.CreateHandArguments
		setMock   func(*mocks)
		want      *domain.Hand
		playerIDs []uint64
		errFunc   func(error) bool
	}{
		{
			name: "success",
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
			setMock: func(m *mocks) {
				c := context.Background()
				halfRoundGameScores := domain.HalfRoundGameScores{
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
				}

				m.playerMock.EXPECT().GetPlayerByID(c, uint64(1)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(2)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(3)).Return(nil, nil)

				m.rdbStatementMock.EXPECT().Transaction(c, gomock.Any()).DoAndReturn(
					func(c context.Context, f func(context.Context) error) error { return f(c) },
				)

				m.handMock.EXPECT().CreateHand(
					c, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				).Return(domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)), nil)
				m.playerHandMock.EXPECT().CreatePlayerHandPairs(c, gomock.Any()).Return(nil)
				m.halfRoundGameMock.EXPECT().CreateHalfRoundGames(
					c, uint64(10), halfRoundGameScores,
				).Return(nil)
			},
			want:      domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
			playerIDs: []uint64{1, 2, 3},
			errFunc:   notErrFunc,
		},
		{
			name: "error in CreateHalfRoundGames",
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
			setMock: func(m *mocks) {
				c := context.Background()
				halfRoundGameScores := domain.HalfRoundGameScores{
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
				}

				m.playerMock.EXPECT().GetPlayerByID(c, uint64(1)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(2)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(3)).Return(nil, nil)

				m.rdbStatementMock.EXPECT().Transaction(c, gomock.Any()).DoAndReturn(
					func(c context.Context, f func(context.Context) error) error { return f(c) },
				)

				m.handMock.EXPECT().CreateHand(
					c, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				).Return(domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)), nil)
				m.playerHandMock.EXPECT().CreatePlayerHandPairs(c, gomock.Any()).Return(nil)
				m.halfRoundGameMock.EXPECT().CreateHalfRoundGames(
					c, uint64(10), halfRoundGameScores,
				).Return(errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "error in CreatePlayerHandPairs",
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
			setMock: func(m *mocks) {
				c := context.Background()

				m.playerMock.EXPECT().GetPlayerByID(c, uint64(1)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(2)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(3)).Return(nil, nil)

				m.rdbStatementMock.EXPECT().Transaction(c, gomock.Any()).DoAndReturn(
					func(c context.Context, f func(context.Context) error) error { return f(c) },
				)

				m.handMock.EXPECT().CreateHand(
					c, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				).Return(domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)), nil)
				m.playerHandMock.EXPECT().CreatePlayerHandPairs(c, gomock.Any()).Return(errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "error in CreateHand",
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
			setMock: func(m *mocks) {
				c := context.Background()

				m.playerMock.EXPECT().GetPlayerByID(c, uint64(1)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(2)).Return(nil, nil)
				m.playerMock.EXPECT().GetPlayerByID(c, uint64(3)).Return(nil, nil)

				m.rdbStatementMock.EXPECT().Transaction(c, gomock.Any()).DoAndReturn(
					func(c context.Context, f func(context.Context) error) error { return f(c) },
				)

				m.handMock.EXPECT().CreateHand(
					c, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				).Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "error in GetPlayerByID",
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
			setMock: func(m *mocks) {
				c := context.Background()

				m.playerMock.EXPECT().GetPlayerByID(c, gomock.Any()).Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "invalid half round game scores",
			args: &usecase.CreateHandArguments{
				Timestamp: time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				PlayerScores: []usecase.CreateHandArgumentsPlayerScore{
					{PlayerID: 1, Score: 10, GameNumber: 1},
					{PlayerID: 1, Score: 0, GameNumber: 2},
					{PlayerID: 2, Score: 20, GameNumber: 1},
					{PlayerID: 3, Score: -30, GameNumber: 1},
					{PlayerID: 3, Score: 5, GameNumber: 3},
					{PlayerID: 3, Score: 40, GameNumber: 2},
					{PlayerID: 2, Score: -40, GameNumber: 2},
					{PlayerID: 2, Score: 10, GameNumber: 3},
					{PlayerID: 1, Score: -10, GameNumber: 3},
				},
			},
			setMock: func(m *mocks) {},
			errFunc: func(err error) bool { return errors.As(err, &domain.InvalidArgumentError{}) },
		},
		{
			name: "empty player scores",
			args: &usecase.CreateHandArguments{
				Timestamp:    time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
				PlayerScores: []usecase.CreateHandArgumentsPlayerScore{},
			},
			setMock: func(m *mocks) {},
			errFunc: func(err error) bool { return errors.As(err, &domain.InvalidArgumentError{}) },
		},
		{
			name:    "empty",
			args:    &usecase.CreateHandArguments{},
			setMock: func(m *mocks) {},
			errFunc: func(err error) bool { return errors.As(err, &domain.InvalidArgumentError{}) },
		},
		{
			name:    "nil",
			args:    nil,
			setMock: func(m *mocks) {},
			errFunc: func(err error) bool { return errors.As(err, &domain.InvalidArgumentError{}) },
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			mocks, finish := setRepository(t)
			defer finish()

			tc.setMock(mocks)
			uc := usecase.NewHandUsecase()

			got, playerIDs, err := uc.CreateHand(context.Background(), tc.args)

			if !tc.errFunc(err) {
				t.Fatalf("unexpected error (error : %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.playerIDs, playerIDs, uintArraySort); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandUsecase_FetchHandScore(t *testing.T) {
	testCases := []struct {
		name      string
		handID    uint64
		setMock   func(*mocks)
		handScore *domain.HandScore
		playerIDs []uint64
		err       bool
	}{
		{
			name:   "success",
			handID: 100,
			setMock: func(m *mocks) {
				c := context.Background()

				m.handMock.EXPECT().GetHandByID(c, uint64(100)).Return(
					domain.NewHand(100, time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(100)).Return([]uint64{1, 2, 3, 9}, nil)
				m.halfRoundGameMock.EXPECT().GetHalfRoundGameScoresByHandID(c, uint64(100)).Return(
					domain.HalfRoundGameScores{
						1: []*domain.PlayerScore{
							domain.NewPlayerScore(1, 10, 2),
							domain.NewPlayerScore(2, 20, 1),
							domain.NewPlayerScore(3, -20, 4),
							domain.NewPlayerScore(9, -10, 3),
						},
					},
					nil,
				)
			},
			handScore: domain.NewHandScore(
				100,
				time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC),
				domain.HalfRoundGameScores{
					1: []*domain.PlayerScore{
						domain.NewPlayerScore(1, 10, 2),
						domain.NewPlayerScore(2, 20, 1),
						domain.NewPlayerScore(3, -20, 4),
						domain.NewPlayerScore(9, -10, 3),
					},
				},
			),
			playerIDs: []uint64{1, 2, 3, 9},
		},
		{
			name:   "error in GetHalfRoundGameScoresByHandID",
			handID: 100,
			setMock: func(m *mocks) {
				c := context.Background()

				m.handMock.EXPECT().GetHandByID(c, uint64(100)).Return(
					domain.NewHand(100, time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(100)).Return([]uint64{1, 2, 3, 9}, nil)
				m.halfRoundGameMock.EXPECT().GetHalfRoundGameScoresByHandID(c, uint64(100)).Return(nil, errors.New("error"))
			},
			err: true,
		},
		{
			name:   "error in ParticipatePlayersInHand",
			handID: 100,
			setMock: func(m *mocks) {
				c := context.Background()

				m.handMock.EXPECT().GetHandByID(c, uint64(100)).Return(
					domain.NewHand(100, time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(100)).Return(nil, errors.New("error"))
			},
			err: true,
		},
		{
			name:   "error in GetHandByID",
			handID: 100,
			setMock: func(m *mocks) {
				c := context.Background()

				m.handMock.EXPECT().GetHandByID(c, uint64(100)).Return(nil, errors.New("error"))
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
			uc := usecase.NewHandUsecase()

			handScore, playerIDs, err := uc.FetchHandScore(context.Background(), tc.handID)

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if diff := cmp.Diff(tc.handScore, handScore, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.playerIDs, playerIDs, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandUsecase_FetcHands(t *testing.T) {
	testCases := []struct {
		name            string
		setMock         func(*mocks)
		hands           []*domain.Hand
		playerIDsInHand map[uint64][]uint64
		err             bool
	}{
		{
			name: "success",
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHands(c).Return(
					[]*domain.Hand{
						domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
						domain.NewHand(20, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
					},
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(10)).Return([]uint64{1, 4, 21}, nil)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(20)).Return([]uint64{3}, nil)
			},
			hands: []*domain.Hand{
				domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
				domain.NewHand(20, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
			},
			playerIDsInHand: map[uint64][]uint64{
				10: {1, 4, 21},
				20: {3},
			},
		},
		{
			name: "hands is emptry",
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHands(c).Return([]*domain.Hand{}, nil)
			},
			hands:           []*domain.Hand{},
			playerIDsInHand: map[uint64][]uint64{},
		},
		{
			name: "error in ParticipatePlayersInHands",
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHands(c).Return(
					[]*domain.Hand{
						domain.NewHand(10, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
						domain.NewHand(20, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
					},
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, gomock.Any()).Return(nil, errors.New("error"))
			},
			err: true,
		},
		{
			name: "error in GetHands",
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHands(c).Return(nil, errors.New("error"))
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
			uc := usecase.NewHandUsecase()

			hands, playerIDsInHand, err := uc.FetchHands(context.Background())

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if diff := cmp.Diff(tc.hands, hands, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.playerIDsInHand, playerIDsInHand); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandUsecase_UpdateHandScore(t *testing.T) {
	testCases := []struct {
		name      string
		args      *usecase.UpdateHandScoreArguments
		setMock   func(*mocks)
		handScore *domain.HandScore
		playerIDs []uint64
		errFunc   func(error) bool
	}{
		{
			name: "success",
			args: &usecase.UpdateHandScoreArguments{
				HandID: 23,
				PlayerScores: map[uint32][]*usecase.UpdateHandScoreArgumentPlayerScore{
					2: {
						{PlayerID: 2, Score: 20},
						{PlayerID: 8, Score: -20},
					},
					12: {
						{PlayerID: 3, Score: -40},
						{PlayerID: 7, Score: -30},
						{PlayerID: 5, Score: 50},
					},
					99: {
						{PlayerID: 3, Score: -40},
						{PlayerID: 7, Score: -30},
						{PlayerID: 5, Score: 50},
					},
				},
			},
			setMock: func(m *mocks) {
				c := context.Background()
				halfRoundGameScores := domain.HalfRoundGameScores{
					1: []*domain.PlayerScore{
						domain.NewPlayerScore(2, 10, 2),
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -10, 3),
						domain.NewPlayerScore(8, -30, 4),
					},
					2: []*domain.PlayerScore{
						domain.NewPlayerScore(2, -10, 3),
						domain.NewPlayerScore(3, -30, 4),
						domain.NewPlayerScore(7, 30, 1),
						domain.NewPlayerScore(8, 10, 2),
					},
					12: []*domain.PlayerScore{
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -30, 4),
						domain.NewPlayerScore(7, -20, 3),
						domain.NewPlayerScore(8, 20, 2),
					},
				}
				m.handMock.EXPECT().GetHandByID(c, uint64(23)).Return(
					domain.NewHand(23, time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(23)).Return([]uint64{2, 3, 5, 7, 8}, nil)
				m.halfRoundGameMock.EXPECT().GetHalfRoundGameScoresByHandID(c, uint64(23)).Return(halfRoundGameScores, nil)

				m.rdbStatementMock.EXPECT().Transaction(c, gomock.Any()).DoAndReturn(
					func(c context.Context, f func(context.Context) error) error { return f(c) },
				)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(2), 20, uint32(2), uint32(2),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(3), -30, uint32(4), uint32(2),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(7), 30, uint32(1), uint32(2),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(8), -20, uint32(3), uint32(2),
				).Return(nil)

				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(3), -40, uint32(4), uint32(12),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(5), 50, uint32(1), uint32(12),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(7), -30, uint32(3), uint32(12),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(8), 20, uint32(2), uint32(12),
				).Return(nil)
			},
			handScore: domain.NewHandScore(
				23,
				time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC),
				domain.HalfRoundGameScores{
					1: []*domain.PlayerScore{
						domain.NewPlayerScore(2, 10, 2),
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -10, 3),
						domain.NewPlayerScore(8, -30, 4),
					},
					2: []*domain.PlayerScore{
						domain.NewPlayerScore(2, 20, 2),
						domain.NewPlayerScore(3, -30, 4),
						domain.NewPlayerScore(7, 30, 1),
						domain.NewPlayerScore(8, -20, 3),
					},
					12: []*domain.PlayerScore{
						domain.NewPlayerScore(3, -40, 4),
						domain.NewPlayerScore(5, 50, 1),
						domain.NewPlayerScore(7, -30, 3),
						domain.NewPlayerScore(8, 20, 2),
					},
				},
			),
			playerIDs: []uint64{2, 3, 5, 7, 8},
			errFunc:   notErrFunc,
		},
		{
			name: "no PlayerScore",
			args: &usecase.UpdateHandScoreArguments{
				HandID:       23,
				PlayerScores: map[uint32][]*usecase.UpdateHandScoreArgumentPlayerScore{},
			},
			setMock: func(m *mocks) {
				c := context.Background()
				halfRoundGameScores := domain.HalfRoundGameScores{
					1: []*domain.PlayerScore{
						domain.NewPlayerScore(2, 10, 2),
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -10, 3),
						domain.NewPlayerScore(8, -30, 4),
					},
					2: []*domain.PlayerScore{
						domain.NewPlayerScore(2, -10, 3),
						domain.NewPlayerScore(3, -30, 4),
						domain.NewPlayerScore(7, 30, 1),
						domain.NewPlayerScore(8, 10, 2),
					},
					12: []*domain.PlayerScore{
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -30, 4),
						domain.NewPlayerScore(7, -20, 3),
						domain.NewPlayerScore(8, 20, 2),
					},
				}
				m.handMock.EXPECT().GetHandByID(c, uint64(23)).Return(
					domain.NewHand(23, time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(23)).Return([]uint64{2, 3, 5, 7, 8}, nil)
				m.halfRoundGameMock.EXPECT().GetHalfRoundGameScoresByHandID(c, uint64(23)).Return(halfRoundGameScores, nil)
			},
			handScore: domain.NewHandScore(
				23,
				time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC),
				domain.HalfRoundGameScores{
					1: []*domain.PlayerScore{
						domain.NewPlayerScore(2, 10, 2),
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -10, 3),
						domain.NewPlayerScore(8, -30, 4),
					},
					2: []*domain.PlayerScore{
						domain.NewPlayerScore(2, -10, 3),
						domain.NewPlayerScore(3, -30, 4),
						domain.NewPlayerScore(7, 30, 1),
						domain.NewPlayerScore(8, 10, 2),
					},
					12: []*domain.PlayerScore{
						domain.NewPlayerScore(3, 30, 1),
						domain.NewPlayerScore(5, -30, 4),
						domain.NewPlayerScore(7, -20, 3),
						domain.NewPlayerScore(8, 20, 2),
					},
				},
			),
			playerIDs: []uint64{2, 3, 5, 7, 8},
			errFunc:   notErrFunc,
		},
		{
			name: "error in UpdateScoreAndRanking",
			args: &usecase.UpdateHandScoreArguments{
				HandID: 23,
				PlayerScores: map[uint32][]*usecase.UpdateHandScoreArgumentPlayerScore{
					2: {
						{PlayerID: 2, Score: 20},
						{PlayerID: 8, Score: -20},
					},
				},
			},
			setMock: func(m *mocks) {
				c := context.Background()
				halfRoundGameScores := domain.HalfRoundGameScores{
					2: []*domain.PlayerScore{
						domain.NewPlayerScore(2, -10, 3),
						domain.NewPlayerScore(3, -30, 4),
						domain.NewPlayerScore(7, 30, 1),
						domain.NewPlayerScore(8, 10, 2),
					},
				}
				m.handMock.EXPECT().GetHandByID(c, uint64(23)).Return(
					domain.NewHand(23, time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(23)).Return([]uint64{2, 3, 5, 7, 8}, nil)
				m.halfRoundGameMock.EXPECT().GetHalfRoundGameScoresByHandID(c, uint64(23)).Return(halfRoundGameScores, nil)

				m.rdbStatementMock.EXPECT().Transaction(c, gomock.Any()).DoAndReturn(
					func(c context.Context, f func(context.Context) error) error { return f(c) },
				)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(2), 20, uint32(2), uint32(2),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(3), -30, uint32(4), uint32(2),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(7), 30, uint32(1), uint32(2),
				).Return(nil)
				m.halfRoundGameMock.EXPECT().UpdateScoreAndRanking(
					c, uint64(23), uint64(8), -20, uint32(3), uint32(2),
				).Return(errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "error in GetHanfScore",
			args: &usecase.UpdateHandScoreArguments{
				HandID: 23,
				PlayerScores: map[uint32][]*usecase.UpdateHandScoreArgumentPlayerScore{
					2: {
						{PlayerID: 2, Score: 20},
						{PlayerID: 8, Score: -20},
					},
				},
			},
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHandByID(c, uint64(23)).Return(
					domain.NewHand(23, time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(23)).Return([]uint64{2, 3, 5, 7, 8}, nil)
				m.halfRoundGameMock.EXPECT().GetHalfRoundGameScoresByHandID(c, uint64(23)).Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "error in GetParticipatePlayerIDs",
			args: &usecase.UpdateHandScoreArguments{
				HandID: 23,
				PlayerScores: map[uint32][]*usecase.UpdateHandScoreArgumentPlayerScore{
					2: {
						{PlayerID: 2, Score: 20},
						{PlayerID: 8, Score: -20},
					},
				},
			},
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHandByID(c, uint64(23)).Return(
					domain.NewHand(23, time.Date(2021, time.November, 22, 0, 0, 0, 0, time.UTC)),
					nil,
				)
				m.playerHandMock.EXPECT().ParticipatePlayersInHand(c, uint64(23)).Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name: "error in GetHandByID",
			args: &usecase.UpdateHandScoreArguments{
				HandID: 23,
				PlayerScores: map[uint32][]*usecase.UpdateHandScoreArgumentPlayerScore{
					2: {
						{PlayerID: 2, Score: 20},
						{PlayerID: 8, Score: -20},
					},
				},
			},
			setMock: func(m *mocks) {
				c := context.Background()
				m.handMock.EXPECT().GetHandByID(c, uint64(23)).Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
		{
			name:    "args is nil",
			setMock: func(m *mocks) {},
			errFunc: errFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			mocks, finish := setRepository(t)
			defer finish()

			tc.setMock(mocks)
			uc := usecase.NewHandUsecase()

			handScore, playerIDs, err := uc.UpdateHandScore(context.Background(), tc.args)

			if !tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.handScore, handScore, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.playerIDs, playerIDs); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
