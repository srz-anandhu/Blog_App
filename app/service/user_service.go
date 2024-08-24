package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
)

type UserService interface {
	GetUser(r *http.Request) (*dto.UserResponse, error)
}

type UserServiceImpl struct {
	userRepo repo.Repo
}

// For checking implementation of UserService interface
var _ UserService = (*UserServiceImpl)(nil)

func NewUserService(userRepo repo.Repo) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) GetUser(r *http.Request) (*dto.UserResponse, error) {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	result, err := s.userRepo.GetOne(req.ID)
	if err != nil {
		return nil, err
	}
	var user dto.UserResponse

	u, ok := result.(repo.User)
	if !ok {
		return nil, err
	}

	user.ID = u.ID
	user.UserName = u.UserName
	user.Password = u.Password
	user.Salt = u.Salt
	user.CreatedAt = u.CreatedAt
	user.UpdatedAt = u.UpdatedAt
	user.DeletedAt = u.DeletedAt
	user.IsDeleted = u.IsDeleted

	return &user, nil
}
