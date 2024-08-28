package e

// 400 errors
const (
	// ErrInvalidRequestGetAuthor : when path param is invalid for get single author
	ErrInvalidRequestGetAuthor = 400001 + iota

	// ErrValidateRequestGetAuthor : when validate AuthorRequest struct
	ErrValidateRequestGetAuthor 
)