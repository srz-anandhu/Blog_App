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

	userMock := new(mocks.UserService)
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
			req := httptest.NewRequest("GET", "/1", nil)
			res := httptest.NewRecorder()

			userMock.On("GetUser", req).Once().Return(test.user, test.err)
			conn.GetUser(res, req)

			assert.Equal(t, test.want, res.Body.String())
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
