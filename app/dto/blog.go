package dto

import (
	"encoding/json"
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

// For Path param
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

// For Body param
type BlogCreateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int    `json:"author_id"` // Author.ID
	Status    int    `json:"status"`
	CreatedBy int    `json:"created_by"` // User.ID
}

func (b *BlogCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}
