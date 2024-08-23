package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"net/http"
)

type BlogController interface {
	GetAllBlogs(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
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
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, blogResp)
}

func (c *blogControllerImpl) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogResp, err := c.blogService.GetAllBlogs()
	if err != nil {
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, blogResp)
}
