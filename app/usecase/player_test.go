package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/usecase"
	"github.com/google/go-cmp/cmp"
)

func TestPlayerUsecase_CreatePlayer(t *testing.T) {
	testCases := []struct {
		testName string
		argName  string
		setMock  func(*mocks)
		want     *domain.Player
		err      bool
	}{
		{
			testName: "create success",
			argName:  "test",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByName(context.Background(), "test").Return(
					nil, domain.NewNotFoundError("error"),
				)
				m.playerMock.EXPECT().CreatePlayer(context.Background(), "test").Return(
					domain.NewPlayer(10, "test"), nil,
				)
			},
			want: domain.NewPlayer(10, "test"),
		},
		{
			testName: "create error",
			argName:  "test",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByName(context.Background(), "test").Return(
					nil, domain.NewNotFoundError("error"),
				)
				m.playerMock.EXPECT().CreatePlayer(context.Background(), "test").Return(
					nil, errors.New("error"),
				)
			},
			err: true,
		},
		{
			testName: "already exist",
			argName:  "test",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByName(context.Background(), "test").Return(
					domain.NewPlayer(10, "test"), nil,
				)
			},
			want: domain.NewPlayer(10, "test"),
		},
		{
			testName: "get error",
			argName:  "test",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByName(context.Background(), "test").Return(
					nil, errors.New("error"),
				)
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			mocks, finish := setRepository(t)
			defer finish()

			tc.setMock(mocks)

			uc := usecase.NewPlayerUsecase()
			got, err := uc.CreatePlayer(context.Background(), tc.argName)

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerUsecase_FetchPlayers(t *testing.T) {
	testCases := []struct {
		name    string
		setMock func(*mocks)
		want    []*domain.Player
		err     bool
	}{
		{
			name: "success",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayers(context.Background()).Return(
					[]*domain.Player{
						domain.NewPlayer(10, "test"),
						domain.NewPlayer(20, "hoge"),
					},
					nil,
				)
			},
			want: []*domain.Player{
				domain.NewPlayer(10, "test"),
				domain.NewPlayer(20, "hoge"),
			},
		},
		{
			name: "empty",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayers(context.Background()).Return([]*domain.Player{}, nil)
			},
			want: []*domain.Player{},
		},
		{
			name: "error",
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayers(context.Background()).Return(nil, errors.New("error"))
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

			uc := usecase.NewPlayerUsecase()
			got, err := uc.FetchPlayers(context.Background())

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerUsecase_UpdatePlayer(t *testing.T) {
	testCases := []struct {
		testName string
		id       uint64
		argName  string
		setMock  func(*mocks)
		want     *domain.Player
		err      bool
	}{
		{
			testName: "success",
			argName:  "test",
			id:       10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(domain.NewPlayer(10, "hoge"), nil)
				m.playerMock.EXPECT().UpdatePlayer(context.Background(), uint64(10), "test").Return(nil)
			},
			want: domain.NewPlayer(10, "test"),
		},
		{
			testName: "same name",
			argName:  "test",
			id:       10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(domain.NewPlayer(10, "test"), nil)
			},
			want: domain.NewPlayer(10, "test"),
		},
		{
			testName: "error in UpdatePlayer",
			argName:  "test",
			id:       10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(domain.NewPlayer(10, "hoge"), nil)
				m.playerMock.EXPECT().UpdatePlayer(context.Background(), uint64(10), "test").Return(errors.New("error"))
			},
			err: true,
		},
		{
			testName: "error in GetPlayerByID",
			argName:  "test",
			id:       10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(nil, errors.New("error"))
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			mocks, finish := setRepository(t)
			defer finish()

			tc.setMock(mocks)

			uc := usecase.NewPlayerUsecase()
			got, err := uc.UpdatePlayer(context.Background(), tc.id, tc.argName)

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerUsecase_DeletePlayer(t *testing.T) {
	testCases := []struct {
		name    string
		id      uint64
		setMock func(*mocks)
		err     bool
	}{
		{
			name: "success",
			id:   10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(domain.NewPlayer(10, "hoge"), nil)
				m.playerMock.EXPECT().DeletePlayer(context.Background(), uint64(10)).Return(nil)
			},
		},
		{
			name: "error in UpdatePlayer",
			id:   10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(domain.NewPlayer(10, "hoge"), nil)
				m.playerMock.EXPECT().DeletePlayer(context.Background(), uint64(10)).Return(errors.New("error"))
			},
			err: true,
		},
		{
			name: "error in GetPlayerByID",
			id:   10,
			setMock: func(m *mocks) {
				m.playerMock.EXPECT().GetPlayerByID(context.Background(), uint64(10)).Return(nil, errors.New("error"))
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

			uc := usecase.NewPlayerUsecase()
			err := uc.DeletePlayer(context.Background(), tc.id)

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}
		})
	}
}
