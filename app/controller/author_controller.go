package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"log"
	"net/http"
)

type AuthorController interface {
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
	CreateAuthor(w http.ResponseWriter, r *http.Request)
}

// For checking implementation of AuthorController interface
var _ AuthorController = (*authorControllerImpl)(nil)

type authorControllerImpl struct {
	authorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &authorControllerImpl{
		authorService: authorService,
	}
}

func (c *authorControllerImpl) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := c.authorService.GetAuthors()
	if err != nil {
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, authors)
}

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {

	authorResponse, err := c.authorService.GetAuthor(r)
	if err != nil {
		log.Printf("can't get author due to : %s", err)
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, authorResponse)
}

func (c *authorControllerImpl) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.DeleteAuthor(r); err != nil {
		api.Fail(w, http.StatusInternalServerError, "author delete failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Deleted author successfully")
}

func (c *authorControllerImpl) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, err := c.authorService.CreateAuthor(r)
	if err != nil {
		api.Fail(w, http.StatusUnprocessableEntity, "failed to create author", err.Error())
		return
	}
	api.Success(w, http.StatusCreated, authorID)
}
