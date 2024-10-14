package service

import (
	"blog/app/dto"
	"blog/app/repo/mocks"
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthor(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	deletedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	authorRepoMock := new(mocks.AuthorRepo)
	conn := NewAuthorService(authorRepoMock)

	tests := []struct {
		name     string
		reqParam string
		args     *dto.AuthorRequest
		author   *dto.AuthorResponse // for mock
		getErr   error
		err      error
		want     *dto.AuthorResponse
		wantErr  bool
	}{
		{
			name:     "Get author service layer success case",
			reqParam: "1",
			args:     &dto.AuthorRequest{ID: 1},
			author: &dto.AuthorResponse{
				ID:   1,
				Name: "demo author name",
				CreateUpdateResponse: dto.CreateUpdateResponse{
					UpdatedAt: &updatedAt,
					UpdatedBy: nil,
					CreatedAt: createdAt,
					CreatedBy: nil,
				},
				DeleteInfoResponse: dto.DeleteInfoResponse{
					DeletedAt: &deletedAt,
					DeletedBy: nil,
				},
			},
			getErr: nil,
			err:    nil,
			want: &dto.AuthorResponse{
				ID:   1,
				Name: "demo author name",
				CreateUpdateResponse: dto.CreateUpdateResponse{
					UpdatedAt: &updatedAt,
					UpdatedBy: nil,
					CreatedAt: createdAt,
					CreatedBy: nil,
				},
				DeleteInfoResponse: dto.DeleteInfoResponse{
					DeletedAt: &deletedAt,
					DeletedBy: nil,
				},
			},
			wantErr: false,
		},
		{
			name:     "Get author service layer error case",
			reqParam: "2",
			args:     &dto.AuthorRequest{ID: 2},
			author:   nil,
			getErr:   errors.New("query failed"),
			// err: e.NewError(e.ErrResourceNotFound, "not found author with requested id", &e.WrapError{
			// 	ErrorCode: 404001,
			// 	Msg:       "not found author with requested id",
			// 	RootCause: errors.New("query failed"),
			// }),
			err:     errors.New("query failed"),
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/", nil)

			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("id", test.reqParam)

			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, chiCtx))
			authorRepoMock.On("GetOne", test.args.ID).Once().Return(test.author, test.err)
			got, err := conn.GetAuthor(request)

			if test.wantErr {
				assert.Equal(t, test.err.Error(), err.Error())
			}
			assert.Equal(t, test.want, got)
		})
	}
}
