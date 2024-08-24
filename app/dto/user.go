package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type UserResponse struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	IsDeleted bool   `json:"is_deleted"`
	CreateUpdateResponse
	DeleteInfoResponse
}

type UserRequest struct {
	ID int `validate:"required"`
}

// For Path param
func (u *UserRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	u.ID = intID
	return nil
}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

// For Body param
type UserCreateRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return err
	}
	return nil
}

func (u *UserCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}
