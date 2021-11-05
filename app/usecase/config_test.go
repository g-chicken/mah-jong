package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	mock_domain "github.com/g-chicken/mah-jong/app/mock/domain"
	"github.com/g-chicken/mah-jong/app/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestConfigUsecase_GetConfig(t *testing.T) {
	testCases := []struct {
		name    string
		setMock func(*mock_domain.MockConfigRepository)
		want    *domain.Config
		err     bool
	}{
		{
			name: "success",
			setMock: func(m *mock_domain.MockConfigRepository) {
				m.EXPECT().GetConfig(context.Background()).Return(
					domain.NewConfig(8080, "url", "name", "user", "pass", 5*time.Second),
					nil,
				)
			},
			want: domain.NewConfig(8080, "url", "name", "user", "pass", 5*time.Second),
		},
		{
			name: "error",
			setMock: func(m *mock_domain.MockConfigRepository) {
				m.EXPECT().GetConfig(context.Background()).Return(nil, errors.New("error"))
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_domain.NewMockConfigRepository(ctrl)
			uc := usecase.NewConfigUsecase(m)

			tc.setMock(m)

			got, err := uc.GetConfig(context.Background())

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
