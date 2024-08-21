package repo

import "time"

type Model struct {
	CreatedBy int // User.ID
	UpdatedBy int // User.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeleteInfo struct {
	DeletedBy int // User.ID
	DeletedAt time.Time
}
