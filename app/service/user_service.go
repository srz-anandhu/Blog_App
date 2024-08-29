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
	UpdateUser(r *http.Request) error
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

	user.ID = result.ID
	user.UserName = result.UserName
	user.Password = result.Password
	user.Salt = result.Salt
	user.CreatedAt = result.CreatedAt
	user.UpdatedAt = result.UpdatedAt
	user.DeletedAt = result.DeletedAt
	user.IsDeleted = result.IsDeleted

	return &user, nil
}

func (s *UserServiceImpl) GetAllUsers() (*[]dto.UserResponse, error) {
	results, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var users []dto.UserResponse

	for _, val := range *results {

		var user dto.UserResponse

		user.ID = val.ID
		user.UserName = val.UserName
		user.Password = val.Password
		user.Salt = val.Salt
		user.CreatedAt = val.CreatedAt
		user.UpdatedAt = val.UpdatedAt
		user.IsDeleted = val.IsDeleted
		user.DeletedAt = val.DeletedAt

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

func (s *UserServiceImpl) UpdateUser(r *http.Request) error {
	body := &dto.UserUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	if err := s.userRepo.Update(body); err != nil {
		return err
	}
	return nil
}
