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
