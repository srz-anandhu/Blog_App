package dto

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type BlogResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   int    `json:"status"` // 1 - Draft, 2 - Published, 3 - Deleted
	AuthorID int    `json:"author_id"`
	CreateUpdateResponse
	DeleteInfoResponse
}

type BlogRequest struct {
	ID int `validate:"required"`
}

func (b *BlogRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	b.ID = intID
	return nil
}

func (b *BlogRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}
