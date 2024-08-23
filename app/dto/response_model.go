package dto

import "time"

type CreateUpdateResponse struct {
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int       `json:"updated_by,omitempty"` // User.ID
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *int       `json:"created_by"` // User.ID
}

type DeleteInfoResponse struct {
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy *int       `json:"deleted_by,omitempty"` // User.ID
}
