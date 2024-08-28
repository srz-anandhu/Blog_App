package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"blog/pkg/e"
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
		httpErr := e.NewAPIError(err, "can't get user")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, userResp)
}

func (c *userControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userResp, err := c.userService.GetAllUsers()
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get all users")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, userResp)
}

func (c *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.DeleteUser(r); err != nil {
		httpErr := e.NewAPIError(err, "user deletion failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Deleted user successfully")
}

func (c *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := c.userService.CreateUser(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "user creation failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusCreated, userID)
}

func (c *userControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.UpdateUser(r); err != nil {
		httpErr := e.NewAPIError(err, "user updation failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "user updated successfully")
}
