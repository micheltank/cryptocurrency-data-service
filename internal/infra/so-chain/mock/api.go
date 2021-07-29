// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock_so_chain is a generated GoMock package.
package mock_so_chain

import (
	so_chain "micheltank/cryptocurrency-data-service/internal/infra/so-chain"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISoChainBlockchainApi is a mock of ISoChainBlockchainApi interface.
type MockISoChainBlockchainApi struct {
	ctrl     *gomock.Controller
	recorder *MockISoChainBlockchainApiMockRecorder
}

// MockISoChainBlockchainApiMockRecorder is the mock recorder for MockISoChainBlockchainApi.
type MockISoChainBlockchainApiMockRecorder struct {
	mock *MockISoChainBlockchainApi
}

// NewMockISoChainBlockchainApi creates a new mock instance.
func NewMockISoChainBlockchainApi(ctrl *gomock.Controller) *MockISoChainBlockchainApi {
	mock := &MockISoChainBlockchainApi{ctrl: ctrl}
	mock.recorder = &MockISoChainBlockchainApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISoChainBlockchainApi) EXPECT() *MockISoChainBlockchainApiMockRecorder {
	return m.recorder
}

// GetBlock mocks base method.
func (m *MockISoChainBlockchainApi) GetBlock(ctx context.Context, networkCode, hash string) (so_chain.BlockResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", ctx, networkCode, hash)
	ret0, _ := ret[0].(so_chain.BlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockISoChainBlockchainApiMockRecorder) GetBlock(ctx, networkCode, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockISoChainBlockchainApi)(nil).GetBlock), ctx, networkCode, hash)
}

// GetTransaction mocks base method.
func (m *MockISoChainBlockchainApi) GetTransaction(ctx context.Context, networkCode, id string) (so_chain.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", ctx, networkCode, id)
	ret0, _ := ret[0].(so_chain.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockISoChainBlockchainApiMockRecorder) GetTransaction(ctx, networkCode, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockISoChainBlockchainApi)(nil).GetTransaction), ctx, networkCode, id)
}