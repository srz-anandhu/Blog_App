package controller

import (
	"blog/app/dto"
	"blog/pkg/api"
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

	var users []dto.UserResponse

	user1 := dto.UserResponse{
		ID:       2,
		UserName: "bbbbb@gmail.com",
	}
	user2 := dto.UserResponse{
		ID:       3,
		UserName: "ccccc@gmail.com",
	}

	users = append(users, user1, user2)

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Printf("error due to : %s ", err)

		api.Fail(w, http.StatusInternalServerError, "failed", "couldn't get users")
		return
	}
	api.Success(w, http.StatusOK, jsonData)
}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	user := dto.UserResponse{
		ID:       1,
		UserName: "aaaa@gmail.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Printf("error due to : %s ", err)

		api.Fail(w, http.StatusInternalServerError, "failed", "couldn't get user")
		return
	}
	api.Success(w, http.StatusOK, jsonData)
}
