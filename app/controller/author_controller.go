package controller

import (
	"blog/app/dto"
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

}

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {
	author := dto.AuthorResponse{
		ID:   1,
		Name: "author bbbb",
	}

	jsonData, err := json.Marshal(author)
	if err != nil {
		log.Printf("error due to : %s ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}
