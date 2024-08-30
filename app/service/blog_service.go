package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"blog/pkg/e"
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
		return nil, e.NewError(e.ErrInvalidRequest, "blog request parse error", err)
	}
	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "blog request validate error", err)
	}
	// Calling GetOne function from repo
	result, err := s.blogRepo.GetOne(req.ID)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "not found blog with requested id", err)
	}

	// Instance of blog response
	var blog dto.BlogResponse

	blog.ID = result.ID
	blog.Title = result.Title
	blog.Content = result.Content
	blog.AuthorID = result.AuthorID
	blog.Status = result.Status
	blog.CreatedBy = result.CreatedBy
	blog.CreatedAt = result.CreatedAt
	blog.UpdatedBy = result.UpdatedBy
	blog.UpdatedAt = result.UpdatedAt
	blog.DeletedBy = result.DeletedBy
	blog.DeletedAt = result.DeletedAt

	return &blog, nil
}

func (s *BlogServiceImpl) GetAllBlogs() (*[]dto.BlogResponse, error) {
	results, err := s.blogRepo.GetAll()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get blogs", err)
	}
	var blogs []dto.BlogResponse

	for _, val := range *results {

		var blogResp dto.BlogResponse

		blogResp.ID = val.ID
		blogResp.Title = val.Title
		blogResp.Content = val.Content
		blogResp.AuthorID = val.AuthorID
		blogResp.Status = val.Status
		blogResp.CreatedBy = val.CreatedBy
		blogResp.CreatedAt = val.CreatedAt
		blogResp.UpdatedBy = val.UpdatedBy
		blogResp.UpdatedAt = val.UpdatedAt
		blogResp.DeletedBy = val.DeletedBy
		blogResp.DeletedAt = val.DeletedAt

		blogs = append(blogs, blogResp)
	}
	return &blogs, nil
}

func (s *BlogServiceImpl) DeleteBlog(r *http.Request) error {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "blog id parse error", err)
	}
	if err := req.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "blog id validation error", err)
	}

	if err := s.blogRepo.Delete(req.ID); err != nil {
		return e.NewError(e.ErrInternalServer, "can't delete blog", err)
	}
	return nil
}

func (s *BlogServiceImpl) CreateBlog(r *http.Request) (int64, error) {
	body := &dto.BlogCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "can't decode blog create request", err)
	}
	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "can't validate blog create request", err)
	}
	blogID, err := s.blogRepo.Create(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create blog", err)
	}
	return blogID, nil
}

func (s *BlogServiceImpl) UpdateBlog(r *http.Request) error {
	body := &dto.BlogUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "can't decode blog update request", err)
	}
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "can't validate blog update request", err)
	}
	if err := s.blogRepo.Update(body); err != nil {
		return e.NewError(e.ErrInternalServer, "can't update blog", err)
	}
	return nil
}
