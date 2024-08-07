package controller

import "net/http"

type AuthorController interface {
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
}

var _ AuthorController = (*authorControllerImpl)(nil)

type authorControllerImpl struct{}

func NewAuthorController()AuthorController{
	return &authorControllerImpl{}
}

func (c *authorControllerImpl)GetAllAuthors(w http.ResponseWriter, r *http.Request){

}

func (c *authorControllerImpl)GetAuthor(w http.ResponseWriter, r *http.Request){

}