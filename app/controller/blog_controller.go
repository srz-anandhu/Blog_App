package controller

import (
	"blog/app/dto"
	"blog/pkg/utils"
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

	var blogs []dto.BlogResponse

	blog1 := dto.BlogResponse{
		ID:      2,
		Title:   "title2",
		Content: "content2",
	}
	blog2 := dto.BlogResponse{
		ID:      3,
		Title:   "title3",
		Content: "content3",
	}

	blogs = append(blogs, blog1, blog2)

	jsonData, err := json.Marshal(blogs)
	if err != nil {
		log.Printf("error due to : %s ", err)
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("failed"))
		utils.Fail(w, http.StatusInternalServerError, []byte("failed"))
		return
	}
	utils.Success(w, http.StatusOK, jsonData)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(jsonData))

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
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("failed"))
		utils.Fail(w, http.StatusInternalServerError, []byte("failed"))
		return
	}
	utils.Success(w, http.StatusOK, jsonData)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(jsonData))
}
