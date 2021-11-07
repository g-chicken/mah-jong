package handler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/g-chicken/mah-jong/app/controller/handler"
	"github.com/g-chicken/mah-jong/app/domain"
	mock_usecase "github.com/g-chicken/mah-jong/app/mock/usecase"
	"github.com/g-chicken/mah-jong/app/proto/app/services/player/v1"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestPlayerServiceServer_CreatePlayer(t *testing.T) {
	testCases := []struct {
		name    string
		req     *player.CreatePlayerRequest
		setMock func(*mock_usecase.MockPlayerUsecase)
		want    *player.CreatePlayerResponse
		errFunc func(error) bool
	}{
		{
			name: "success",
			req:  &player.CreatePlayerRequest{Name: "test"},
			setMock: func(m *mock_usecase.MockPlayerUsecase) {
				m.EXPECT().CreatePlayer(context.Background(), "test").Return(
					domain.NewPlayer(10, "test"), nil,
				)
			},
			want: &player.CreatePlayerResponse{
				Player: &player.Player{
					Id:   10,
					Name: "test",
				},
			},
			errFunc: noErrFunc,
		},
		{
			name:    "no name",
			req:     &player.CreatePlayerRequest{Name: ""},
			setMock: func(_ *mock_usecase.MockPlayerUsecase) {},
			errFunc: func(err error) bool { return !errors.As(err, &domain.InvalidArgumentError{}) },
		},
		{
			name: "error",
			req:  &player.CreatePlayerRequest{Name: "test"},
			setMock: func(m *mock_usecase.MockPlayerUsecase) {
				m.EXPECT().CreatePlayer(context.Background(), "test").Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_usecase.NewMockPlayerUsecase(ctrl)
			tc.setMock(m)
			service := handler.NewPlayerServiceServer(m)

			got, err := service.CreatePlayer(context.Background(), tc.req)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (err : %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, ignoreUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerServiceServer_FetchPlayers(t *testing.T) {
	testCases := []struct {
		name    string
		req     *player.FetchPlayersRequest
		setMock func(*mock_usecase.MockPlayerUsecase)
		want    *player.FetchPlayersResponse
		errFunc func(err error) bool
	}{
		{
			name: "success",
			req:  &player.FetchPlayersRequest{},
			setMock: func(m *mock_usecase.MockPlayerUsecase) {
				m.EXPECT().FetchPlayers(context.Background()).Return(
					[]*domain.Player{
						domain.NewPlayer(10, "test"),
						domain.NewPlayer(20, "hoge"),
					},
					nil,
				)
			},
			want: &player.FetchPlayersResponse{
				Players: []*player.Player{
					{Id: 10, Name: "test"},
					{Id: 20, Name: "hoge"},
				},
			},
			errFunc: noErrFunc,
		},
		{
			name: "empty",
			req:  &player.FetchPlayersRequest{},
			setMock: func(m *mock_usecase.MockPlayerUsecase) {
				m.EXPECT().FetchPlayers(context.Background()).Return(
					[]*domain.Player{},
					nil,
				)
			},
			want: &player.FetchPlayersResponse{
				Players: []*player.Player{},
			},
			errFunc: noErrFunc,
		},
		{
			name: "error",
			req:  &player.FetchPlayersRequest{},
			setMock: func(m *mock_usecase.MockPlayerUsecase) {
				m.EXPECT().FetchPlayers(context.Background()).Return(nil, errors.New("error"))
			},
			errFunc: errFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_usecase.NewMockPlayerUsecase(ctrl)
			tc.setMock(m)
			service := handler.NewPlayerServiceServer(m)

			got, err := service.FetchPlayers(context.Background(), tc.req)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (err : %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, ignoreUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
