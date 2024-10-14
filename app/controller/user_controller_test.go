package controller

import (
	"blog/app/dto"
	"blog/app/service/mocks"
	"blog/pkg/e"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	deletedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)

	// Creating an instance of user service mock
	userMock := new(mocks.UserService)
	// Passing mocked service layer(user) to NewUserController to get a connection
	conn := NewUserController(userMock)

	tests := []struct {
		name    string
		status  int
		want    string
		user    *dto.UserResponse
		err     error
		wantErr bool
	}{
		{
			name:   "get user success case",
			status: 200,
			want:   `{"status":"ok","result":{"id":4,"username":"random@gmail.com","password":"randompassword","salt":"1234asdf","is_deleted":false,"updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"}}`,
			user: &dto.UserResponse{
				ID:        4,
				UserName:  "random@gmail.com",
				Password:  "randompassword",
				Salt:      "1234asdf",
				IsDeleted: false,
				CreateUpdateResponse: dto.CreateUpdateResponse{
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
					UpdatedBy: nil,
					CreatedBy: nil,
				},
				DeleteInfoResponse: dto.DeleteInfoResponse{
					DeletedAt: &deletedAt,
					DeletedBy: nil,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:   "get user error case",
			status: 400,
			want:   `{"status":"not ok","error":{"code":400,"message":"can't get user","details":["Bad Request"]}}`,
			err: &e.WrapError{
				ErrorCode: 400,
				Msg:       "can't get user",
				RootCause: errors.New("Bad Request"),
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// creating a dummy request
			req := httptest.NewRequest("GET", "/1", nil)
			res := httptest.NewRecorder()
			// calling "GetUser" function from user service
			userMock.On("GetUser", req).Once().Return(test.user, test.err)
			// calling "GetUser" function from user controller
			conn.GetUser(res, req)
			// checking want data with response
			assert.Equal(t, test.want, res.Body.String())
			// checking status code with response code
			assert.Equal(t, test.status, res.Code)
		})
	}
}

func TestCreateUser(t *testing.T) {
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)

	tests := []struct {
		name       string
		status     int
		want       string
		userCreate *dto.UserCreateRequest
		userID     int64
		err        error
		wantErr    bool
	}{
		{
			name:   "create user success case",
			status: 201,
			want:   `{"status":"ok","result":7}`,
			userCreate: &dto.UserCreateRequest{
				UserName: "demouser name",
				Password: "demouser password",
			},
			userID:  7,
			err:     nil,
			wantErr: false,
		},

		{
			name:   "create user error case",
			status: 500,
			want:   `{"status":"not ok","error":{"code":500,"message":"user creation failed","details":["Internal server error"]}}`,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "user creation failed",
				RootCause: errors.New("Internal server error"),
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/create", nil)
			res := httptest.NewRecorder()

			// calling service layer function
			userMock.On("CreateUser", req).Once().Return(test.userID, test.err)
			// calling controller layer function
			conn.CreateUser(res, req)
			// checking wanted data with response
			assert.Equal(t, test.want, res.Body.String())
			// checking status code with response code
			assert.Equal(t, test.status, res.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)

	tests := []struct {
		name    string
		status  int
		want    string
		err     error
		wantErr bool
	}{
		{
			name:    "delete user success case",
			status:  200,
			want:    `{"status":"ok","result":"Deleted user successfully"}`,
			err:     nil,
			wantErr: false,
		},

		{
			name:   "delete user error case",
			status: 500,
			want:   `{"status":"not ok","error":{"code":500,"message":"user deletion failed","details":["Internal server error"]}}`,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "user deletion failed",
				RootCause: errors.New("Internal server error"),
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/4", nil)
			res := httptest.NewRecorder()

			userMock.On("DeleteUser", req).Once().Return(test.err)
			conn.DeleteUser(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)

	tests := []struct {
		name       string
		status     int
		want       string
		userUpdate dto.UserUpdateRequest
		err        error
		wantErr    bool
	}{
		{
			name:   "update user success case",
			status: 200,
			want:   `{"status":"ok","result":"user updated successfully"}`,
			userUpdate: dto.UserUpdateRequest{
				ID:       5,
				Username: "updated user name",
				Password: "somerandom password",
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:   "update user error case",
			status: 500,
			want:   `{"status":"not ok","error":{"code":500,"message":"user updation failed","details":["Internal server error"]}}`,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "user updation failed",
				RootCause: errors.New("Internal server error"),
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/5", nil)
			res := httptest.NewRecorder()

			userMock.On("UpdateUser", req).Once().Return(test.err)
			conn.UpdateUser(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	userMock := new(mocks.UserService)
	conn := NewUserController(userMock)

	tests := []struct {
		name    string
		status  int
		want    string
		users   *[]dto.UserResponse
		err     error
		wantErr bool
	}{
		{
			name:   "get all users success case",
			status: 200,
			want:   `{"status":"ok","result":[{"id":1,"username":"something@gmail.com","password":"some password","salt":"asdfg","is_deleted":false,"updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"},{"id":2,"username":"random@gmail.com","password":"randompassword","salt":"adafda","is_deleted":false,"updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"}]}`,
			users: &[]dto.UserResponse{
				{
					ID:        1,
					UserName:  "something@gmail.com",
					Password:  "some password",
					Salt:      "asdfg",
					IsDeleted: false,
					CreateUpdateResponse: dto.CreateUpdateResponse{
						UpdatedAt: &updatedAt,
						UpdatedBy: nil,
						CreatedAt: createdAt,
						CreatedBy: nil,
					},
					DeleteInfoResponse: dto.DeleteInfoResponse{
						DeletedAt: &createdAt,
						DeletedBy: nil,
					},
				},
				{
					ID:        2,
					UserName:  "random@gmail.com",
					Password:  "randompassword",
					Salt:      "adafda",
					IsDeleted: false,
					CreateUpdateResponse: dto.CreateUpdateResponse{
						UpdatedAt: &updatedAt,
						UpdatedBy: nil,
						CreatedAt: createdAt,
						CreatedBy: nil,
					},
					DeleteInfoResponse: dto.DeleteInfoResponse{
						DeletedAt: &createdAt,
						DeletedBy: nil,
					},
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:   "get all users error case",
			status: 500,
			want:   `{"status":"not ok","error":{"code":500,"message":"can't get all users","details":["Internal server error"]}}`,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "can't get all users",
				RootCause: errors.New("Internal server error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			res := httptest.NewRecorder()

			userMock.On("GetAllUsers").Once().Return(test.users, test.err)
			conn.GetAllUsers(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}
