package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"blog/pkg/e"
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
		return nil, e.NewError(e.ErrInvalidRequest, "author request parse error", err)
	}

	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "author request validate error", err)
	}
	result, err := s.authorRepo.GetOne(req.ID)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "not found author with requested id", err)
	}

	var authorResp dto.AuthorResponse

	authorResp.ID = result.ID
	authorResp.Name = result.Name
	authorResp.CreatedAt = result.CreatedAt
	authorResp.CreatedBy = result.CreatedBy
	authorResp.UpdatedAt = result.UpdatedAt
	authorResp.UpdatedBy = result.UpdatedBy
	authorResp.DeletedAt = result.DeletedAt

	return &authorResp, nil

}

func (s *AuthorServiceImpl) GetAuthors() (*[]dto.AuthorResponse, error) {

	results, err := s.authorRepo.GetAll()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get authors", err)
	}
	var authors []dto.AuthorResponse
	for _, val := range *results {

		var author dto.AuthorResponse
		author.ID = val.ID
		author.Name = val.Name
		author.CreatedAt = val.CreatedAt
		author.CreatedBy = val.CreatedBy
		author.UpdatedAt = val.UpdatedAt
		author.UpdatedBy = val.UpdatedBy

		authors = append(authors, author)
	}
	return &authors, nil

}

func (s *AuthorServiceImpl) DeleteAuthor(r *http.Request) error {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "author id parse error", err)
	}
	if err := req.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "author id validation error", err)
	}
	if err := s.authorRepo.Delete(req.ID); err != nil {
		return e.NewError(e.ErrResourceNotFound, "can't delete author", err)
	}
	return nil
}

func (s *AuthorServiceImpl) CreateAuthor(r *http.Request) (int64, error) {

	// Creating instance of dto.AuthorCreateRequest
	body := &dto.AuthorCreateRequest{}

	// Decode to dto.AuthorCreateRequest
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "can't decode author create request", err)
	}
	// Validating dto.AuthorCreateRequest
	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "can't validate author create request", err)
	}

	authorID, err := s.authorRepo.Create(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create author", err)
	}
	return authorID, nil
}

func (s *AuthorServiceImpl) UpdateAuthor(r *http.Request) error {

	// Creating instance of dto.AuthorUpdateRequest
	body := &dto.AuthorUpdateRequest{}

	// Decode to dto.AuthorUpdateRequest
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrInternalServer, "can't decode author update request", err)
	}
	// Validating dto.AuthorUpdateRequest
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrInternalServer, "can't validate author update request", err)
	}
	if err := s.authorRepo.Update(body); err != nil {
		return e.NewError(e.ErrInternalServer, "can't update author", err)
	}
	return nil
}
