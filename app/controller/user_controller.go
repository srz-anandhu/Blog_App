package controller

import (
	"blog/app/dto"
	"encoding/json"
	"log"
	"net/http"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

var _ UserController = (*userControllerImpl)(nil)

type userControllerImpl struct{}

func NewUserController() UserController {
	return &userControllerImpl{}
}

func (c *userControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	user := dto.UserResponse{
		ID:       1,
		UserName: "aaaa@gmail.com",
	}

	jsonData, err := json.Marshal(user)
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
