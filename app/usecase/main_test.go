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
	playerMock        *mock_domain.MockPlayerRepository
	handMock          *mock_domain.MockHandRepository
	halfRoundGameMock *mock_domain.MockHalfRoundGameRepository
	playerHandMock    *mock_domain.MockPlayerHandRepository
	rdbStatementMock  *mock_domain.MockRDBStatementSetRepository
}

func setRepository(t *testing.T) (*mocks, func()) {
	t.Helper()

	ctrl := gomock.NewController(t)
	playerMock := mock_domain.NewMockPlayerRepository(ctrl)
	handMock := mock_domain.NewMockHandRepository(ctrl)
	halfRoundGameMock := mock_domain.NewMockHalfRoundGameRepository(ctrl)
	playerHandMock := mock_domain.NewMockPlayerHandRepository(ctrl)
	rdbStatementSetMock := mock_domain.NewMockRDBStatementSetRepository(ctrl)

	domain.SetRepositories(
		playerMock,
		handMock,
		halfRoundGameMock,
		playerHandMock,
		rdbStatementSetMock,
	)

	return &mocks{
		playerMock:        playerMock,
		handMock:          handMock,
		halfRoundGameMock: halfRoundGameMock,
		playerHandMock:    playerHandMock,
		rdbStatementMock:  rdbStatementSetMock,
	}, ctrl.Finish
}
