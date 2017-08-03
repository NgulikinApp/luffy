// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

import user "github.com/NgulikinApp/luffy/user"

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *UserRepository) GetByID(id int64) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(int64) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
