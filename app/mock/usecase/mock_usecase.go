// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "github.com/g-chicken/mah-jong/app/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockConfigUsecase is a mock of ConfigUsecase interface.
type MockConfigUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockConfigUsecaseMockRecorder
}

// MockConfigUsecaseMockRecorder is the mock recorder for MockConfigUsecase.
type MockConfigUsecaseMockRecorder struct {
	mock *MockConfigUsecase
}

// NewMockConfigUsecase creates a new mock instance.
func NewMockConfigUsecase(ctrl *gomock.Controller) *MockConfigUsecase {
	mock := &MockConfigUsecase{ctrl: ctrl}
	mock.recorder = &MockConfigUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigUsecase) EXPECT() *MockConfigUsecaseMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockConfigUsecase) GetConfig(c context.Context) (*domain.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", c)
	ret0, _ := ret[0].(*domain.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockConfigUsecaseMockRecorder) GetConfig(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockConfigUsecase)(nil).GetConfig), c)
}

// MockPlayerUsecase is a mock of PlayerUsecase interface.
type MockPlayerUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPlayerUsecaseMockRecorder
}

// MockPlayerUsecaseMockRecorder is the mock recorder for MockPlayerUsecase.
type MockPlayerUsecaseMockRecorder struct {
	mock *MockPlayerUsecase
}

// NewMockPlayerUsecase creates a new mock instance.
func NewMockPlayerUsecase(ctrl *gomock.Controller) *MockPlayerUsecase {
	mock := &MockPlayerUsecase{ctrl: ctrl}
	mock.recorder = &MockPlayerUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayerUsecase) EXPECT() *MockPlayerUsecaseMockRecorder {
	return m.recorder
}

// CreatePlayer mocks base method.
func (m *MockPlayerUsecase) CreatePlayer(c context.Context, name string) (*domain.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlayer", c, name)
	ret0, _ := ret[0].(*domain.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlayer indicates an expected call of CreatePlayer.
func (mr *MockPlayerUsecaseMockRecorder) CreatePlayer(c, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlayer", reflect.TypeOf((*MockPlayerUsecase)(nil).CreatePlayer), c, name)
}
