package repo

import "time"

type Model struct {
	CreatedBy *int       `json:"created_by"` // User.ID
	UpdatedBy *int       `json:"updated_by"` // User.ID
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteInfo struct {
	DeletedBy *int       `json:"deleted_by,omitempty"` // User.ID
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
