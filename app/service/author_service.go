package service

import (
	"blog/app/dto"
	"blog/app/repo"
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
	//fmt.Printf("author id from service : %d", id)
	result, err := s.authorRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	a, ok := result.(repo.Author)
	if !ok {
		return nil, err
	}

	//fmt.Println("assertion before: ", a)
	authRes := &dto.AuthorResponse{
		ID:         a.ID,
		Name:       a.Name,
		Created_at: a.CreatedAt,
		Updated_at: a.UpdatedAt,
	}
	//fmt.Println("::::::", authRes)
	return authRes, nil

}
