// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/interfaces/user_usecase_intf.go

// Package mocksUseCase is a generated GoMock package.
package mocksUseCase

import (
	reflect "reflect"

	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	gomock "github.com/golang/mock/gomock"
)

// MockIuserUseCase is a mock of IuserUseCase interface.
type MockIuserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIuserUseCaseMockRecorder
}

// MockIuserUseCaseMockRecorder is the mock recorder for MockIuserUseCase.
type MockIuserUseCaseMockRecorder struct {
	mock *MockIuserUseCase
}

// NewMockIuserUseCase creates a new mock instance.
func NewMockIuserUseCase(ctrl *gomock.Controller) *MockIuserUseCase {
	mock := &MockIuserUseCase{ctrl: ctrl}
	mock.recorder = &MockIuserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIuserUseCase) EXPECT() *MockIuserUseCaseMockRecorder {
	return m.recorder
}

// ChangeUserStatusById mocks base method.
func (m *MockIuserUseCase) ChangeUserStatusById(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserStatusById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserStatusById indicates an expected call of ChangeUserStatusById.
func (mr *MockIuserUseCaseMockRecorder) ChangeUserStatusById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserStatusById", reflect.TypeOf((*MockIuserUseCase)(nil).ChangeUserStatusById), arg0, arg1)
}

// EditUserData mocks base method.
func (m *MockIuserUseCase) EditUserData(arg0 *requestmodels.UserEditProf) (*responsemodels.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUserData", arg0)
	ret0, _ := ret[0].(*responsemodels.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditUserData indicates an expected call of EditUserData.
func (mr *MockIuserUseCaseMockRecorder) EditUserData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUserData", reflect.TypeOf((*MockIuserUseCase)(nil).EditUserData), arg0)
}

// GetLatestUsers mocks base method.
func (m *MockIuserUseCase) GetLatestUsers() (*[]responsemodels.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestUsers")
	ret0, _ := ret[0].(*[]responsemodels.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestUsers indicates an expected call of GetLatestUsers.
func (mr *MockIuserUseCaseMockRecorder) GetLatestUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestUsers", reflect.TypeOf((*MockIuserUseCase)(nil).GetLatestUsers))
}

// SearchUserByIdOrName mocks base method.
func (m *MockIuserUseCase) SearchUserByIdOrName(arg0 int, arg1 string) (*[]responsemodels.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUserByIdOrName", arg0, arg1)
	ret0, _ := ret[0].(*[]responsemodels.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUserByIdOrName indicates an expected call of SearchUserByIdOrName.
func (mr *MockIuserUseCaseMockRecorder) SearchUserByIdOrName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUserByIdOrName", reflect.TypeOf((*MockIuserUseCase)(nil).SearchUserByIdOrName), arg0, arg1)
}

// UserByStatus mocks base method.
func (m *MockIuserUseCase) UserByStatus(arg0 string) (*[]responsemodels.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByStatus", arg0)
	ret0, _ := ret[0].(*[]responsemodels.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByStatus indicates an expected call of UserByStatus.
func (mr *MockIuserUseCaseMockRecorder) UserByStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByStatus", reflect.TypeOf((*MockIuserUseCase)(nil).UserByStatus), arg0)
}

// UserLogin mocks base method.
func (m *MockIuserUseCase) UserLogin(arg0 *requestmodels.UserLoginReq) (responsemodels.UserLoginRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", arg0)
	ret0, _ := ret[0].(responsemodels.UserLoginRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockIuserUseCaseMockRecorder) UserLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockIuserUseCase)(nil).UserLogin), arg0)
}

// UserProfile mocks base method.
func (m *MockIuserUseCase) UserProfile(arg0 *string) (*responsemodels.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserProfile", arg0)
	ret0, _ := ret[0].(*responsemodels.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserProfile indicates an expected call of UserProfile.
func (mr *MockIuserUseCaseMockRecorder) UserProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserProfile", reflect.TypeOf((*MockIuserUseCase)(nil).UserProfile), arg0)
}

// UserSignUp mocks base method.
func (m *MockIuserUseCase) UserSignUp(arg0 *requestmodels.UserSignUpReq) (responsemodels.SignupData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", arg0)
	ret0, _ := ret[0].(responsemodels.SignupData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockIuserUseCaseMockRecorder) UserSignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockIuserUseCase)(nil).UserSignUp), arg0)
}

// VerifyOtp mocks base method.
func (m *MockIuserUseCase) VerifyOtp(arg0 *requestmodels.OtpVerification, arg1 string) (responsemodels.OtpVerifResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOtp", arg0, arg1)
	ret0, _ := ret[0].(responsemodels.OtpVerifResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyOtp indicates an expected call of VerifyOtp.
func (mr *MockIuserUseCaseMockRecorder) VerifyOtp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOtp", reflect.TypeOf((*MockIuserUseCase)(nil).VerifyOtp), arg0, arg1)
}
