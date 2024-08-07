package controller

import "net/http"

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

var _ UserController = (*userControllerImpl)(nil)

type userControllerImpl struct{}

func NewUserController()UserController{
	return &userControllerImpl{}
}

func (c *userControllerImpl)GetAllUsers(w http.ResponseWriter, r *http.Request){

}

func (c *userControllerImpl)GetUser(w http.ResponseWriter, r *http.Request){
	
}