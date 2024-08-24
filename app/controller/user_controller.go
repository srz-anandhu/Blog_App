package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"net/http"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

// For checking implementation of UserController interface
var _ UserController = (*userControllerImpl)(nil)

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController() UserController {
	return &userControllerImpl{}
}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	userResp, err := c.userService.GetUser(r)
	if err != nil {
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, userResp)
}

func (c *userControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
