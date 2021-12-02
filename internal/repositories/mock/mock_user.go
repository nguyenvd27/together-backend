// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"
	models "together-backend/internal/models"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUserRepo) ChangePassword(user *models.User, hashPassword []byte) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", user, hashPassword)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepoMockRecorder) ChangePassword(user, hashPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepo)(nil).ChangePassword), user, hashPassword)
}

// CreateUser mocks base method.
func (m *MockUserRepo) CreateUser(name, email string, password []byte) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", name, email, password)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoMockRecorder) CreateUser(name, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepo)(nil).CreateUser), name, email, password)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepo) GetUserByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepoMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepo)(nil).GetUserByEmail), email)
}

// GetUserById mocks base method.
func (m *MockUserRepo) GetUserById(id int64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserRepoMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepo)(nil).GetUserById), id)
}

// GetUserDetail mocks base method.
func (m *MockUserRepo) GetUserDetail(id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetail", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetail indicates an expected call of GetUserDetail.
func (mr *MockUserRepoMockRecorder) GetUserDetail(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetail", reflect.TypeOf((*MockUserRepo)(nil).GetUserDetail), id)
}

// UpdateProfile mocks base method.
func (m *MockUserRepo) UpdateProfile(user *models.User, name string, address int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", user, name, address)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserRepoMockRecorder) UpdateProfile(user, name, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserRepo)(nil).UpdateProfile), user, name, address)
}

// UpdateProfileWithAvatar mocks base method.
func (m *MockUserRepo) UpdateProfileWithAvatar(user *models.User, name string, address int, avatarUrl string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileWithAvatar", user, name, address, avatarUrl)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfileWithAvatar indicates an expected call of UpdateProfileWithAvatar.
func (mr *MockUserRepoMockRecorder) UpdateProfileWithAvatar(user, name, address, avatarUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileWithAvatar", reflect.TypeOf((*MockUserRepo)(nil).UpdateProfileWithAvatar), user, name, address, avatarUrl)
}
