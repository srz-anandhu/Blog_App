package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"errors"
)

type AuthorService interface {
	GetAuthor(id int) (*dto.AuthorResponse, error)
	// GetAuthors()
}

var _ AuthorService = (*AuthorServiceImpl)(nil)

type AuthorServiceImpl struct {
	authorRepo repo.Repo
}

func NewAuthorService(authorRepo repo.Repo) AuthorService {
	return &AuthorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *AuthorServiceImpl) GetAuthor(id int) (*dto.AuthorResponse, error) {
	result, err := s.authorRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	a, ok := result.(*dto.AuthorResponse)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	authRes := &dto.AuthorResponse{
		ID:         a.ID,
		Name:       a.Name,
		Created_at: a.Created_at,
		Updated_at: a.Updated_at,
	}

	return authRes, nil

}
