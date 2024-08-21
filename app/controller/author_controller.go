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

type authorControllerImpl struct{
	authorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &authorControllerImpl{
		authorService: authorService,
	}
}

func (c *authorControllerImpl) GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	// var authors []dto.AuthorResponse

	// author1 := dto.AuthorResponse{
	// 	ID:   2,
	// 	Name: "Author Name 8888",
	// }
	// author2 := dto.AuthorResponse{
	// 	ID:   3,
	// 	Name: "Author Name 77777",
	// }

	// authors = append(authors, author1, author2)

	// jsonData, err := json.Marshal(authors)
	// if err != nil {
	// 	log.Printf("error due to : %s ", err)

	// 	api.Fail(w, http.StatusInternalServerError, "failed", "couldn't get authors")
	// 	return
	// }
	// api.Success(w, http.StatusOK, jsonData)
}

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {

	// get author ID from request
	strID := chi.URLParam(r, "id")

	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)
	if err != nil{
		log.Printf("can't get author ID from request due to : %s", err)
		api.Fail(w, http.StatusBadRequest, "failed", "can't found author ID")
		return
	}
	authorResponse, err := c.authorService.GetAuthor(intID)
	if err != nil{
		log.Printf("can't get author due to : %s", err)
		api.Fail(w, http.StatusBadRequest, "failed", "couldn't get author")
		return
	}
	api.Success(w, http.StatusOK, authorResponse)
}
