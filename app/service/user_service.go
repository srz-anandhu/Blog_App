package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
)

type UserService interface {
	GetUser(r *http.Request) (*dto.UserResponse, error)
	GetAllUsers() (*[]dto.UserResponse, error)
	DeleteUser(r *http.Request) error
	CreateUser(r *http.Request) (int64, error)
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

// For checking implementation of UserService interface
var _ UserService = (*UserServiceImpl)(nil)

func NewUserService(userRepo repo.UserRepo) UserService {
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

func (s *UserServiceImpl) GetAllUsers() (*[]dto.UserResponse, error) {
	results, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var users []dto.UserResponse

	for _, val := range results {
		u, ok := val.(repo.User)
		if !ok {
			return nil, err
		}
		var user dto.UserResponse

		user.ID = u.ID
		user.UserName = u.UserName
		user.Password = u.Password
		user.Salt = u.Salt
		user.CreatedAt = u.CreatedAt
		user.UpdatedAt = u.UpdatedAt
		user.IsDeleted = u.IsDeleted
		user.DeletedAt = u.DeletedAt

		users = append(users, user)
	}
	return &users, nil
}

func (s *UserServiceImpl) DeleteUser(r *http.Request) error {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}
	if err := s.userRepo.Delete(req.ID); err != nil {
		return err
	}
	return nil
}

func (c *UserServiceImpl) CreateUser(r *http.Request) (int64, error) {
	body := &dto.UserCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, err
	}
	if err := body.Validate(); err != nil {
		return 0, err
	}
	userID, err := c.userRepo.Create(body)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
