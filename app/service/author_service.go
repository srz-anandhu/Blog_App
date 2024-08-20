package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"database/sql"
	"errors"
)

type AuthorService interface {
	GetAuthor() (*dto.AuthorResponse, error)
	// GetAuthors()
}

var _ AuthorService = (*AuthorServiceImpl)(nil)

type AuthorServiceImpl struct {
	db         *sql.DB
	authorRepo repo.Repo
}

func NewAuthorService(authorRepo repo.Repo, db *sql.DB) AuthorService {
	return &AuthorServiceImpl{
		authorRepo: authorRepo,
		db:         db,
	}
}

func (s *AuthorServiceImpl) GetAuthor() (*dto.AuthorResponse, error) {
	result, err := s.authorRepo.GetOne(s.db)
	if err != nil {
		return nil, err
	}

	a, ok := result.(dto.AuthorResponse)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	var authRes dto.AuthorResponse
	
	authRes.ID = a.ID
	authRes.Name = a.Name
	authRes.Created_at = a.Created_at
	authRes.Updated_at = a.Updated_at

	return &authRes, nil

}
