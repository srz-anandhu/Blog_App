package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"blog/pkg/e"
	"net/http"
)

type AuthorController interface {
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
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
		httpErr := e.NewAPIError(err, "can't get all authors")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, authors)
}

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {

	authorResponse, err := c.authorService.GetAuthor(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, authorResponse)
}

func (c *authorControllerImpl) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.DeleteAuthor(r); err != nil {
		httpErr := e.NewAPIError(err, "author deletion failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Deleted author successfully")
}

func (c *authorControllerImpl) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, err := c.authorService.CreateAuthor(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "failed to create author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusCreated, authorID)
}

func (c *authorControllerImpl) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.UpdateAuthor(r); err != nil {
		httpErr := e.NewAPIError(err, "can't update author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Author updated successfully")
}
