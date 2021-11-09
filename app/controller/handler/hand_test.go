package handler_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/controller/handler"
	"github.com/g-chicken/mah-jong/app/domain"
	mock_usecase "github.com/g-chicken/mah-jong/app/mock/usecase"
	"github.com/g-chicken/mah-jong/app/proto/app/services/hand/v1"
	"github.com/g-chicken/mah-jong/app/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHandServiceServer_CreateHand(t *testing.T) {
	testCases := []struct {
		name    string
		req     *hand.CreateHandRequest
		setMock func(*mock_usecase.MockHandUsecase)
		want    *hand.CreateHandResponse
		errFunc func(error) bool
	}{
		{
			name: "success",
			req: &hand.CreateHandRequest{
				Timestamp: timestamppb.New(time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
				PlayerScores: []*hand.CreateHandRequest_PlayerScore{
					{PlayerId: 10, Score: 10, GameNumber: 1},
					{PlayerId: 11, Score: 5, GameNumber: 1},
					{PlayerId: 12, Score: -15, GameNumber: 1},
				},
			},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				args := &usecase.CreateHandArguments{
					Timestamp: time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
					PlayerScores: []usecase.CreateHandArgumentsPlayerScore{
						{PlayerID: 10, Score: 10, GameNumber: 1},
						{PlayerID: 11, Score: 5, GameNumber: 1},
						{PlayerID: 12, Score: -15, GameNumber: 1},
					},
				}

				m.EXPECT().CreateHand(context.Background(), args).Return(
					domain.NewHand(2, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
					[]uint64{10, 11, 12},
					nil,
				)
			},
			want: &hand.CreateHandResponse{
				Hand: &hand.Hand{
					Id:                   2,
					ParticipatePlayerIds: []uint64{10, 11, 12},
					Timestamp:            timestamppb.New(time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
				},
			},
			errFunc: noErrFunc,
		},
		{
			name: "error",
			req: &hand.CreateHandRequest{
				Timestamp: timestamppb.New(time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
				PlayerScores: []*hand.CreateHandRequest_PlayerScore{
					{PlayerId: 10, Score: 10, GameNumber: 1},
					{PlayerId: 11, Score: 5, GameNumber: 1},
					{PlayerId: 12, Score: -15, GameNumber: 1},
				},
			},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				args := &usecase.CreateHandArguments{
					Timestamp: time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC),
					PlayerScores: []usecase.CreateHandArgumentsPlayerScore{
						{PlayerID: 10, Score: 10, GameNumber: 1},
						{PlayerID: 11, Score: 5, GameNumber: 1},
						{PlayerID: 12, Score: -15, GameNumber: 1},
					},
				}

				m.EXPECT().CreateHand(context.Background(), args).Return(nil, nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_usecase.NewMockHandUsecase(ctrl)
			tc.setMock(m)
			service := handler.NewHandServiceServer(m)

			got, err := service.CreateHand(context.Background(), tc.req)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (err : %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, ignoreUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandServiceServer_FetchHandScore(t *testing.T) {
	testCases := []struct {
		name    string
		req     *hand.FetchHandScoreRequest
		setMock func(*mock_usecase.MockHandUsecase)
		want    *hand.FetchHandScoreResponse
		errFunc func(error) bool
	}{
		{
			name: "success",
			req:  &hand.FetchHandScoreRequest{Id: 100},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				m.EXPECT().FetchHandScore(context.Background(), uint64(100)).Return(
					domain.NewHand(100, time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					[]uint64{3, 8, 11},
					domain.HalfRoundGameScores{
						1: []*domain.PlayerScore{
							domain.NewPlayerScore(3, -21, 3),
							domain.NewPlayerScore(8, -9, 2),
							domain.NewPlayerScore(11, 30, 1),
						},
						2: []*domain.PlayerScore{
							domain.NewPlayerScore(3, 30, 1),
							domain.NewPlayerScore(8, -9, 2),
							domain.NewPlayerScore(11, -21, 3),
						},
					},
					nil,
				)
			},
			want: &hand.FetchHandScoreResponse{
				HandScore: &hand.HandScore{
					Id:                   100,
					ParticipatePlayerIds: []uint64{3, 8, 11},
					Timestamp:            timestamppb.New(time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					HalfGameScores: map[uint32]*hand.HandScore_HalfGameScore{
						1: {
							PlayerScores: []*hand.HandScore_HalfGameScore_PlayerScore{
								{PlayerId: 3, Score: -21, Ranking: 3},
								{PlayerId: 8, Score: -9, Ranking: 2},
								{PlayerId: 11, Score: 30, Ranking: 1},
							},
						},
						2: {
							PlayerScores: []*hand.HandScore_HalfGameScore_PlayerScore{
								{PlayerId: 3, Score: 30, Ranking: 1},
								{PlayerId: 8, Score: -9, Ranking: 2},
								{PlayerId: 11, Score: -21, Ranking: 3},
							},
						},
					},
				},
			},
			errFunc: noErrFunc,
		},
		{
			name: "empty scores",
			req:  &hand.FetchHandScoreRequest{Id: 100},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				m.EXPECT().FetchHandScore(context.Background(), uint64(100)).Return(
					domain.NewHand(100, time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					[]uint64{3, 8, 11},
					domain.HalfRoundGameScores{},
					nil,
				)
			},
			want: &hand.FetchHandScoreResponse{
				HandScore: &hand.HandScore{
					Id:                   100,
					ParticipatePlayerIds: []uint64{3, 8, 11},
					Timestamp:            timestamppb.New(time.Date(2021, time.November, 9, 0, 0, 0, 0, time.UTC)),
					HalfGameScores:       map[uint32]*hand.HandScore_HalfGameScore{},
				},
			},
			errFunc: noErrFunc,
		},
		{
			name: "error",
			req:  &hand.FetchHandScoreRequest{Id: 100},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				m.EXPECT().FetchHandScore(context.Background(), uint64(100)).Return(nil, nil, nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_usecase.NewMockHandUsecase(ctrl)
			tc.setMock(m)
			service := handler.NewHandServiceServer(m)

			got, err := service.FetchHandScore(context.Background(), tc.req)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (err : %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, ignoreUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandServiceServer_FetchHands(t *testing.T) {
	testCases := []struct {
		name    string
		req     *hand.FetchHandsRequest
		setMock func(*mock_usecase.MockHandUsecase)
		want    *hand.FetchHandsResponse
		errFunc func(error) bool
	}{
		{
			name: "success",
			req:  &hand.FetchHandsRequest{},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				m.EXPECT().FetchHands(context.Background()).Return(
					[]*domain.Hand{
						domain.NewHand(2, time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
						domain.NewHand(3, time.Date(2021, time.November, 8, 0, 0, 0, 0, time.UTC)),
					},
					map[uint64][]uint64{
						2: {10, 11, 12},
						3: {2},
					},
					nil,
				)
			},
			want: &hand.FetchHandsResponse{
				Hands: []*hand.Hand{
					{
						Id:                   2,
						ParticipatePlayerIds: []uint64{10, 11, 12},
						Timestamp:            timestamppb.New(time.Date(2021, time.November, 7, 0, 0, 0, 0, time.UTC)),
					},
					{
						Id:                   3,
						ParticipatePlayerIds: []uint64{2},
						Timestamp:            timestamppb.New(time.Date(2021, time.November, 8, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			errFunc: noErrFunc,
		},
		{
			name: "error",
			req:  &hand.FetchHandsRequest{},
			setMock: func(m *mock_usecase.MockHandUsecase) {
				m.EXPECT().FetchHands(context.Background()).Return(nil, nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_usecase.NewMockHandUsecase(ctrl)
			tc.setMock(m)
			service := handler.NewHandServiceServer(m)

			got, err := service.FetchHands(context.Background(), tc.req)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (err : %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, ignoreUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
