package service

import (
	"blog/app/dto"
	"net/http"
)

type BlogService interface {
	GetBlog(r *http.Request) (*dto.BlogResponse)
}