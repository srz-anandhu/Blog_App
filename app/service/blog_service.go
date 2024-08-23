package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BlogService interface {
	GetBlog(r *http.Request) (*dto.BlogResponse, error)
	GetAllBlogs() (*[]dto.BlogResponse, error)
}

type BlogServiceImpl struct {
	blogRepo repo.Repo
}

// For checking implementation of BlogService interface
var _ BlogService = (*BlogServiceImpl)(nil)

func NewBlogService(blogRepo repo.Repo) BlogService {
	return &BlogServiceImpl{
		blogRepo: blogRepo,
	}
}

func (s *BlogServiceImpl) GetBlog(r *http.Request) (*dto.BlogResponse, error) {

	// Get Blog ID from request
	strID := chi.URLParam(r, "id")
	// Converting to Int
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return nil, err
	}

	// Calling GetOne function from repo
	result, err := s.blogRepo.GetOne(intID)
	if err != nil {
		return nil, err
	}
	// type assertion
	b, ok := result.(repo.Blog)
	if !ok {
		return nil, err
	}
	// Instance of blog response
	var blog dto.BlogResponse

	blog.ID = b.ID
	blog.Title = b.Title
	blog.Content = b.Content
	blog.AuthorID = b.AuthorID
	blog.CreatedBy = b.CreatedBy
	blog.CreatedAt = b.CreatedAt
	blog.UpdatedBy = b.UpdatedBy
	blog.UpdatedAt = b.UpdatedAt
	blog.DeletedBy = b.UpdatedBy
	blog.DeletedAt = b.DeletedAt

	return &blog, nil
}

func (s *BlogServiceImpl) GetAllBlogs() (*[]dto.BlogResponse, error) {
	results, err := s.blogRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var blogs []dto.BlogResponse

	for _, val := range results {
		b, ok := val.(repo.Blog)
		if !ok {
			return nil, err
		}
		var blogResp dto.BlogResponse

		blogResp.ID = b.ID
		blogResp.Title = b.Title
		blogResp.Content = b.Content
		blogResp.AuthorID = b.AuthorID
		blogResp.CreatedBy = b.CreatedBy
		blogResp.CreatedAt = b.CreatedAt
		blogResp.UpdatedBy = b.UpdatedBy
		blogResp.UpdatedAt = b.UpdatedAt

		blogs = append(blogs, blogResp)
	}
	return &blogs, nil
}
