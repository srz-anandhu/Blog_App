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

func TestGetBlog(t *testing.T) {
	createdAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	deletedAt := time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)
	blogMock := new(mocks.BlogService)
	conn := NewBlogController(blogMock)

	tests := []struct {
		name    string
		status  int
		want    string
		blog    *dto.BlogResponse
		err     error
		wantErr bool
	}{
		{
			name:   "Get blog success case",
			status: 200,
			want:   `{"status":"ok","result":{"id":3,"title":"some blog title","content":"some blog content","status":2,"author_id":1,"updated_at":"2024-07-15T00:00:00Z","created_at":"2024-07-15T00:00:00Z","deleted_at":"2024-07-15T00:00:00Z"}}`,
			blog: &dto.BlogResponse{
				ID:       3,
				Title:    "some blog title",
				Content:  "some blog content",
				Status:   2,
				AuthorID: 1,
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
			err:     nil,
			wantErr: false,
		},
		{
			name:   "Get blog error case",
			status: 404, // resource not found
			want:   `{"status":"not ok","error":{"code":404,"message":"failed to get blog","details":["Resource not found"]}}`,
			err: &e.WrapError{
				ErrorCode: 404,
				Msg:       "failed to get blog",
				RootCause: errors.New("Resource not found"),
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/3", nil)
			res := httptest.NewRecorder()

			blogMock.On("GetBlog", req).Once().Return(test.blog, test.err)
			conn.GetBlog(res, req)

			assert.Equal(t, test.want, res.Body.String())
			assert.Equal(t, test.status, res.Code)
		})
	}
}
