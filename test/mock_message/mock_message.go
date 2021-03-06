// Code generated by MockGen. DO NOT EDIT.
// Source: messageservice.go

// Package mock_message is a generated GoMock package.
package mock_message

import (
	reflect "reflect"

	common "talos-sdk-golang/thrift/common"
	x "talos-sdk-golang/thrift/message"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageService is a mock of MessageService interface
type MockMessageService struct {
	ctrl     *gomock.Controller
	recorder *MockMessageServiceMockRecorder
}

// MockMessageServiceMockRecorder is the mock recorder for MockMessageService
type MockMessageServiceMockRecorder struct {
	mock *MockMessageService
}

// NewMockMessageService creates a new mock instance
func NewMockMessageService(ctrl *gomock.Controller) *MockMessageService {
	mock := &MockMessageService{ctrl: ctrl}
	mock.recorder = &MockMessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageService) EXPECT() *MockMessageServiceMockRecorder {
	return m.recorder
}

// GetServiceVersion mocks base method
func (m *MockMessageService) GetServiceVersion() (*common.Version, error) {
	ret := m.ctrl.Call(m, "GetServiceVersion")
	ret0, _ := ret[0].(*common.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceVersion indicates an expected call of GetServiceVersion
func (mr *MockMessageServiceMockRecorder) GetServiceVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceVersion", reflect.TypeOf((*MockMessageService)(nil).GetServiceVersion))
}

// ValidClientVersion mocks base method
func (m *MockMessageService) ValidClientVersion(clientVersion *common.Version) error {
	ret := m.ctrl.Call(m, "ValidClientVersion", clientVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidClientVersion indicates an expected call of ValidClientVersion
func (mr *MockMessageServiceMockRecorder) ValidClientVersion(clientVersion interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidClientVersion", reflect.TypeOf((*MockMessageService)(nil).ValidClientVersion), clientVersion)
}

// PutMessage mocks base method
func (m *MockMessageService) PutMessage(request *x.PutMessageRequest) (*x.PutMessageResponse, error) {
	ret := m.ctrl.Call(m, "PutMessage", request)
	ret0, _ := ret[0].(*x.PutMessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutMessage indicates an expected call of PutMessage
func (mr *MockMessageServiceMockRecorder) PutMessage(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMessage", reflect.TypeOf((*MockMessageService)(nil).PutMessage), request)
}

// GetMessage mocks base method
func (m *MockMessageService) GetMessage(request *x.GetMessageRequest) (*x.GetMessageResponse, error) {
	ret := m.ctrl.Call(m, "GetMessage", request)
	ret0, _ := ret[0].(*x.GetMessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage
func (mr *MockMessageServiceMockRecorder) GetMessage(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockMessageService)(nil).GetMessage), request)
}

// GetTopicOffset mocks base method
func (m *MockMessageService) GetTopicOffset(request *x.GetTopicOffsetRequest) (*x.GetTopicOffsetResponse, error) {
	ret := m.ctrl.Call(m, "GetTopicOffset", request)
	ret0, _ := ret[0].(*x.GetTopicOffsetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicOffset indicates an expected call of GetTopicOffset
func (mr *MockMessageServiceMockRecorder) GetTopicOffset(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicOffset", reflect.TypeOf((*MockMessageService)(nil).GetTopicOffset), request)
}

// GetPartitionOffset mocks base method
func (m *MockMessageService) GetPartitionOffset(request *x.GetPartitionOffsetRequest) (*x.GetPartitionOffsetResponse, error) {
	ret := m.ctrl.Call(m, "GetPartitionOffset", request)
	ret0, _ := ret[0].(*x.GetPartitionOffsetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionOffset indicates an expected call of GetPartitionOffset
func (mr *MockMessageServiceMockRecorder) GetPartitionOffset(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionOffset", reflect.TypeOf((*MockMessageService)(nil).GetPartitionOffset), request)
}

// GetPartitionsOffset mocks base method
func (m *MockMessageService) GetPartitionsOffset(request *x.GetPartitionsOffsetRequest) (*x.GetPartitionsOffsetResponse, error) {
	ret := m.ctrl.Call(m, "GetPartitionsOffset", request)
	ret0, _ := ret[0].(*x.GetPartitionsOffsetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitionsOffset indicates an expected call of GetPartitionsOffset
func (mr *MockMessageServiceMockRecorder) GetPartitionsOffset(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionsOffset", reflect.TypeOf((*MockMessageService)(nil).GetPartitionsOffset), request)
}

// GetScheduleInfo mocks base method
func (m *MockMessageService) GetScheduleInfo(request *x.GetScheduleInfoRequest) (*x.GetScheduleInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetScheduleInfo", request)
	ret0, _ := ret[0].(*x.GetScheduleInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduleInfo indicates an expected call of GetScheduleInfo
func (mr *MockMessageServiceMockRecorder) GetScheduleInfo(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleInfo", reflect.TypeOf((*MockMessageService)(nil).GetScheduleInfo), request)
}
