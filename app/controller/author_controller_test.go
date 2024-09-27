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

func TestGetAuthor(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)
	tests := []struct {
		name    string
		status  int
		want    string // dto.AuthorResponse
		author  *dto.AuthorResponse
		error   error
		wantErr bool
	}{
		// Success case
		{
			name:   "success case",
			status: 200,
			author: &dto.AuthorResponse{
				ID:   1,
				Name: "testing name",
				CreateUpdateResponse: dto.CreateUpdateResponse{
					CreatedAt: createdAt,
					CreatedBy: nil,
					UpdatedAt: &updatedAt,
					UpdatedBy: nil,
				},
				DeleteInfoResponse: dto.DeleteInfoResponse{
					DeletedAt: &createdAt,
					DeletedBy: nil,
				},
			},
			want:    `{"status":"ok","result":{"id":1,"name":"testing name","updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"}}`,
			error:   nil,
			wantErr: false,
		},

		// Error case
		{
			name:   "error case",
			status: 400,
			error: &e.WrapError{
				ErrorCode: 400,
				Msg:       "Bad request",
				RootCause: errors.New("invalid request"),
			},
			want:    `{"status":"not ok","error":{"code":400,"message":"can't get author","details":["invalid request"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/1", nil)
			res := httptest.NewRecorder()
			authorMock.On("GetAuthor", req).Once().Return(test.author, test.error)
			conn.GetAuthor(res, req)

			assert.Equal(t, test.status, res.Code)
			assert.Equal(t, test.want, res.Body.String())
		})
	}
}

func TestDeleteAuthor(t *testing.T) {
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)
	tests := []struct {
		name    string
		status  int
		want    string
		err   error
		wantErr bool
	}{
		// success case
		{
			name:    "delete author success case",
			status:  200,
			want:    `{"status":"ok","result":"Deleted author successfully"}`,
			err:   nil,
			wantErr: false,
		},

		// error case
		{
			name:   "error case",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "Internal server error",
				RootCause: errors.New("Internal server error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"author deletion failed","details":["Internal server error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/1", nil)
			res := httptest.NewRecorder()
			authorMock.On("DeleteAuthor", req).Once().Return(test.err)
			conn.DeleteAuthor(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}

func TestCreateAuthor(t *testing.T) {
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)

	tests := []struct {
		name         string
		status       int
		authorCreate *dto.AuthorCreateRequest
		authorID     int64
		want         string
		err          error
		wantErr      bool
	}{
		// success case
		{
			name:   "author creation success case",
			status: 201,
			authorCreate: &dto.AuthorCreateRequest{
				Name:      "newauthorname",
				CreatedBy: 1, // User ID
			},
			authorID: 2,
			want:     `{"status":"ok","result":2}`,
			err:      nil,
			wantErr:  false,
		},
		// error case
		{
			name:   "author creation error case",
			status: 500,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg:       "failed to create author",
				RootCause: errors.New("Internal server error"),
			},
			want:    `{"status":"not ok","error":{"code":500,"message":"failed to create author","details":["Internal server error"]}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/create", nil)
			res := httptest.NewRecorder()
			authorMock.On("CreateAuthor", req).Once().Return(test.authorID, test.err)
			conn.CreateAuthor(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}

func TestGetAllAuthors(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	authorMock := new(mocks.AuthorService)
	conn := NewAuthorController(authorMock)

	tests := []struct {
		name    string
		status  int
		authors *[]dto.AuthorResponse
		want    string
		err     error
		wantErr bool
	}{
		// success case
		{
			name:   "Get all authors success case",
			status: 200,
			authors: &[]dto.AuthorResponse{
				{
					ID:   1,
					Name: "author name1",
					CreateUpdateResponse: dto.CreateUpdateResponse{
						CreatedAt: createdAt,
						CreatedBy: nil,
						UpdatedAt: &updatedAt,
						UpdatedBy: nil,
					},
					DeleteInfoResponse: dto.DeleteInfoResponse{
						DeletedAt: &createdAt,
						DeletedBy: nil,
					},
				},
				{
					ID:   2,
					Name: "author name2",
					CreateUpdateResponse: dto.CreateUpdateResponse{
						CreatedAt: createdAt,
						CreatedBy: nil,
						UpdatedAt: &updatedAt,
						UpdatedBy: nil,
					},
					DeleteInfoResponse: dto.DeleteInfoResponse{
						DeletedAt: &createdAt,
						DeletedBy: nil,
					},
				},
				{
					ID:   3,
					Name: "author name3",
					CreateUpdateResponse: dto.CreateUpdateResponse{
						CreatedAt: createdAt,
						CreatedBy: nil,
						UpdatedAt: &updatedAt,
						UpdatedBy: nil,
					},
					DeleteInfoResponse: dto.DeleteInfoResponse{
						DeletedAt: &createdAt,
						DeletedBy: nil,
					},
				},
			},
			want:    `{"status":"ok","result":[{"id":1,"name":"author name1","updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"},{"id":2,"name":"author name2","updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"},{"id":3,"name":"author name3","updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"}]}`,
			err:     nil,
			wantErr: false,
		},

		// error case
		{
			name: "Get all authors error case",
			status: 500,
			want: `{"status":"not ok","error":{"code":500,"message":"can't get all authors","details":["Internal server error"]}}`,
			err: &e.WrapError{
				ErrorCode: 500,
				Msg: "can't get all authors",
				RootCause: errors.New("Internal server error"),
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			res := httptest.NewRecorder()

			authorMock.On("GetAuthors").Once().Return(test.authors, test.err)
			conn.GetAllAuthors(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}
