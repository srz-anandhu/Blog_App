package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
)

type AuthorService interface {
	GetAuthor(r *http.Request) (*dto.AuthorResponse, error)
	GetAuthors() (*[]dto.AuthorResponse, error)
	DeleteAuthor(r *http.Request) error
	CreateAuthor(r *http.Request) (int64, error)
	UpdateAuthor(r *http.Request) error
}

var _ AuthorService = (*AuthorServiceImpl)(nil)

type AuthorServiceImpl struct {
	authorRepo repo.AuthorRepo
}

func NewAuthorService(authorRepo repo.AuthorRepo) AuthorService {
	return &AuthorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *AuthorServiceImpl) GetAuthor(r *http.Request) (*dto.AuthorResponse, error) {

	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}
	result, err := s.authorRepo.GetOne(req.ID)
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
	authorResp.UpdatedBy = a.UpdatedBy
	authorResp.DeletedAt = a.DeletedAt

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
		author.UpdatedBy = a.UpdatedBy

		authors = append(authors, author)
	}
	return &authors, nil

}

func (s *AuthorServiceImpl) DeleteAuthor(r *http.Request) error {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}
	if err := s.authorRepo.Delete(req.ID); err != nil {
		return err
	}
	return nil
}

func (s *AuthorServiceImpl) CreateAuthor(r *http.Request) (int64, error) {
	body := &dto.AuthorCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, err
	}
	if err := body.Validate(); err != nil {
		return 0, err
	}

	authorID, err := s.authorRepo.Create(body)
	if err != nil {
		return 0, err
	}
	return authorID, nil
}

func (s *AuthorServiceImpl) UpdateAuthor(r *http.Request) error {
	body := &dto.AuthorUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	if err := s.authorRepo.Update(body); err != nil {
		return err
	}
	return nil
}
