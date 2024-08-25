package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
)

type BlogService interface {
	GetBlog(r *http.Request) (*dto.BlogResponse, error)
	GetAllBlogs() (*[]dto.BlogResponse, error)
	DeleteBlog(r *http.Request) error
	CreateBlog(r *http.Request) (int64, error)
	UpdateBlog(r *http.Request) error
}

type BlogServiceImpl struct {
	blogRepo repo.BlogRepo
}

// For checking implementation of BlogService interface
var _ BlogService = (*BlogServiceImpl)(nil)

func NewBlogService(blogRepo repo.BlogRepo) BlogService {
	return &BlogServiceImpl{
		blogRepo: blogRepo,
	}
}

func (s *BlogServiceImpl) GetBlog(r *http.Request) (*dto.BlogResponse, error) {

	req := &dto.BlogRequest{}
	if err := req.Parse(r); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// Calling GetOne function from repo
	result, err := s.blogRepo.GetOne(req.ID)
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
	blog.Status = b.Status
	blog.CreatedBy = b.CreatedBy
	blog.CreatedAt = b.CreatedAt
	blog.UpdatedBy = b.UpdatedBy
	blog.UpdatedAt = b.UpdatedAt
	blog.DeletedBy = b.DeletedBy
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
		blogResp.Status = b.Status
		blogResp.CreatedBy = b.CreatedBy
		blogResp.CreatedAt = b.CreatedAt
		blogResp.UpdatedBy = b.UpdatedBy
		blogResp.UpdatedAt = b.UpdatedAt
		blogResp.DeletedBy = b.DeletedBy
		blogResp.DeletedAt = b.DeletedAt

		blogs = append(blogs, blogResp)
	}
	return &blogs, nil
}

func (s *BlogServiceImpl) DeleteBlog(r *http.Request) error {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}
	if err := s.blogRepo.Delete(req.ID); err != nil {
		return err
	}
	return nil
}

func (s *BlogServiceImpl) CreateBlog(r *http.Request) (int64, error) {
	body := &dto.BlogCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, err
	}
	if err := body.Validate(); err != nil {
		return 0, err
	}
	blogID, err := s.blogRepo.Create(body)
	if err != nil {
		return 0, err
	}
	return blogID, nil
}

func (s *BlogServiceImpl) UpdateBlog(r *http.Request) error {
	body := &dto.BlogUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	if err := s.blogRepo.Update(body); err != nil {
		return err
	}
	return nil
}
