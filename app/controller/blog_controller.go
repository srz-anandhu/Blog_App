package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"blog/pkg/e"
	"net/http"
)

type BlogController interface {
	GetAllBlogs(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
	DeleteBlog(w http.ResponseWriter, r *http.Request)
	CreateBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
}

// For checking implementation of BlogController interface
var _ BlogController = (*blogControllerImpl)(nil)

type blogControllerImpl struct {
	blogService service.BlogService
}

func NewBlogController(blogService service.BlogService) BlogController {
	return &blogControllerImpl{
		blogService: blogService,
	}
}

func (c *blogControllerImpl) GetBlog(w http.ResponseWriter, r *http.Request) {
	blogResp, err := c.blogService.GetBlog(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "failed to get blog")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, blogResp)
}

func (c *blogControllerImpl) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogResp, err := c.blogService.GetAllBlogs()
	if err != nil {
		httpErr := e.NewAPIError(err, "failed to get all blogs")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, blogResp)
}

func (c *blogControllerImpl) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if err := c.blogService.DeleteBlog(r); err != nil {

		httpErr := e.NewAPIError(err, "failed to delete blog")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())

		return
	}
	api.Success(w, http.StatusOK, "Blog deleted successfully")
}

func (c *blogControllerImpl) CreateBlog(w http.ResponseWriter, r *http.Request) {
	blogID, err := c.blogService.CreateBlog(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "blog creation failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusCreated, blogID)
}

func (c *blogControllerImpl) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	if err := c.blogService.UpdateBlog(r); err != nil {
		httpErr := e.NewAPIError(err, "blog updation failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "blog updated successfully")
}
