package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AuthorService interface {
	GetAuthor(r *http.Request) (*dto.AuthorResponse, error)
	GetAuthors() (*[]dto.AuthorResponse, error)
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

func (s *AuthorServiceImpl) GetAuthor(r *http.Request) (*dto.AuthorResponse, error) {

	// get author ID from request
	strID := chi.URLParam(r, "id")
	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)

	if err != nil {
		return nil, err
	}
	result, err := s.authorRepo.GetOne(intID)
	if err != nil {
		return nil, err
	}

	a, ok := result.(repo.Author)
	if !ok {
		return nil, err
	}

	var authorResp dto.AuthorResponse

	authorResp.ID = a.ID
	authorResp.Name = a.Name
	authorResp.CreatedAt = a.CreatedAt
	authorResp.CreatedBy = a.CreatedBy
	authorResp.UpdatedAt = a.UpdatedAt

	return &authorResp, nil

}

func (s *AuthorServiceImpl) GetAuthors() (*[]dto.AuthorResponse, error) {

	results, err := s.authorRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var authors []dto.AuthorResponse
	for _, val := range results {
		a, ok := val.(repo.Author)
		if !ok {
			return nil, err
		}
		var author dto.AuthorResponse
		author.ID = a.ID
		author.Name = a.Name
		author.CreatedAt = a.CreatedAt
		author.CreatedBy = a.CreatedBy
		author.UpdatedAt = a.UpdatedAt

		authors = append(authors, author)
	}
	return &authors, nil

}
