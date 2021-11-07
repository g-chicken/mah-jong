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
