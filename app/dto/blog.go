package dto

import "time"

type BlogResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   int    `json:"status"` // 1 - Draft, 2 - Published, 3 - Deleted
	AuthorID int    `json:"author_id"`
	BlogCreateUpdateResponse
	BlogDeleteInfo
}

type BlogCreateUpdateResponse struct {
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"` // User.ID
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"` // User.ID
}

type BlogDeleteInfo struct {
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy int       `json:"deleted_by,omitempty"` // User.ID
}
