package service

import (
	"blog/app/dto"
	"blog/app/repo/mocks"
	"context"
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/", nil)
			// response := httptest.NewRecorder()
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("id", test.reqParam)

			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, chiCtx))
			authorRepoMock.On("GetOne", test.args.ID).Once().Return(test.author, test.err)
			got, err := conn.GetAuthor(request)

			if test.wantErr {
				assert.Equal(t, test.err, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}
