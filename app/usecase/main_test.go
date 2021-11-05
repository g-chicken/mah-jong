package usecase_test

import (
	"os"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	mock_domain "github.com/g-chicken/mah-jong/app/mock/domain"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var allowUnexported = cmp.AllowUnexported(domain.Config{}, domain.Player{})

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

type mocks struct {
	playerMock *mock_domain.MockPlayerRepository
}

func setRepository(t *testing.T) (*mocks, func()) {
	t.Helper()

	ctrl := gomock.NewController(t)
	playerMock := mock_domain.NewMockPlayerRepository(ctrl)

	domain.SetRepositories(playerMock)

	return &mocks{
		playerMock: playerMock,
	}, ctrl.Finish
}
