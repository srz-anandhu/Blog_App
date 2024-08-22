package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AuthorController interface {
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
}

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
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, authors)
}

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {

	// get author ID from request
	strID := chi.URLParam(r, "id")

	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)
	//fmt.Printf("author id is :: %d ", intID)

	if err != nil {
		log.Printf("can't get author ID from request due to : %s", err)
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}
	authorResponse, err := c.authorService.GetAuthor(intID)
	if err != nil {
		log.Printf("can't get author due to : %s", err)
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, authorResponse)
}
