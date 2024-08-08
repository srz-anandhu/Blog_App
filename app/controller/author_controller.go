package controller

import (
	"blog/app/dto"
	"blog/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
)

type AuthorController interface {
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
}

var _ AuthorController = (*authorControllerImpl)(nil)

type authorControllerImpl struct{}

func NewAuthorController() AuthorController {
	return &authorControllerImpl{}
}

func (c *authorControllerImpl) GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	var authors []dto.AuthorResponse

	author1 := dto.AuthorResponse{
		ID:   2,
		Name: "Author Name 8888",
	}
	author2 := dto.AuthorResponse{
		ID:   3,
		Name: "Author Name 77777",
	}

	authors = append(authors, author1, author2)

	jsonData, err := json.Marshal(authors)
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

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {
	author := dto.AuthorResponse{
		ID:   1,
		Name: "author bbbb",
	}

	jsonData, err := json.Marshal(author)
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
