package dto

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type AuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CreateUpdateResponse
	DeleteInfoResponse
}

type AuthorRequest struct {
	ID int `validate:"required"`
}

func (a *AuthorRequest) Parse(r *http.Request) error {
	// get author ID from request
	strID := chi.URLParam(r, "id")
	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = intID
	return nil
}

func (a *AuthorRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}
