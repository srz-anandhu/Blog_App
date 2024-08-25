package controller

import (
	"blog/app/service"
	"blog/pkg/api"
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

func (c *blogControllerImpl) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if err := c.blogService.DeleteBlog(r); err != nil {
		api.Fail(w, http.StatusInternalServerError, "failed to delete blog", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Blog deleted successfully")
}

func (c *blogControllerImpl) CreateBlog(w http.ResponseWriter, r *http.Request) {
	blogID, err := c.blogService.CreateBlog(r)
	if err != nil {
		api.Fail(w, http.StatusBadRequest, "blog creation failed", err.Error())
		return
	}
	api.Success(w, http.StatusCreated, blogID)
}

func (c *blogControllerImpl) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	if err := c.blogService.UpdateBlog(r); err != nil {
		api.Fail(w, http.StatusInternalServerError, "blog updation failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "blog updated successfully")
}
