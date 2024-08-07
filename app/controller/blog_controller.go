package controller

import (
	"blog/app/dto"
	"encoding/json"
	"log"
	"net/http"
)

type BlogController interface {
	GetAllBlogs(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
}

var _ BlogController = (*blogControllerImpl)(nil)

type blogControllerImpl struct{}

func NewBlogController() BlogController {
	return &blogControllerImpl{}
}

func (c *blogControllerImpl) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("got all blogs"))
}

func (c *blogControllerImpl) GetBlog(w http.ResponseWriter, r *http.Request) {
	blog := dto.BlogResponse{
		ID:      1,
		Title:   "sometitle",
		Content: "some content",
	}

	jsonData, err := json.Marshal(blog)
	if err != nil {
		log.Printf("error due to : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}
