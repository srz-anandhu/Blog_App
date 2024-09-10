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

// For Body param
type BlogCreateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int    `json:"author_id" validate:"required"` // Author.ID
	Status    int    `json:"status"`
	CreatedBy int    `json:"created_by" validate:"required"` // User.ID
}

type BlogUpdateRequest struct {
	ID        int    `validate:"required"`
	Status    int    `validate:"required"` // Update only if status 1 or 2 (drafted or published), 3 is 'deleted'
	Title     string `json:"title"`
	Content   string `json:"content"`
	UpdatedBy int    `json:"updated_by" validate:"required"`
}

type BlogDeleteRequest struct {
	ID        int `json:"id"`
	DeletedBy int `json:"deleted_by"`
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

func (b *BlogUpdateRequest) Parse(r *http.Request) error {
	// Get ID from request
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	b.ID = intID
	// Decode to BlogUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogDeleteRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogDeleteRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}
