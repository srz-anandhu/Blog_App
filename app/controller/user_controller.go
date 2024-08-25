package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"net/http"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

// For checking implementation of UserController interface
var _ UserController = (*userControllerImpl)(nil)

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
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
	userResp, err := c.userService.GetAllUsers()
	if err != nil {
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, userResp)
}

func (c *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.DeleteUser(r); err != nil {
		api.Fail(w, http.StatusInternalServerError, "failed to delete user", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Deleted user successfully")
}

func (c *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := c.userService.CreateUser(r)
	if err != nil {
		api.Fail(w, http.StatusBadRequest, "user creation failed", err.Error())
		return
	}
	api.Success(w, http.StatusCreated, userID)
}

func (c *userControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.UpdateUser(r); err != nil {
		api.Fail(w, http.StatusInternalServerError, "user updation failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "user updated successfully")
}
