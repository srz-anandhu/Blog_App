package dto

type BlogResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   int    `json:"status"` // 1 - Draft, 2 - Published, 3 - Deleted
	AuthorID int    `json:"author_id"`
	CreateUpdateResponse
	DeleteInfoResponse
}
