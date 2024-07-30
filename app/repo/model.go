package repo

import "time"

type Model struct {
	CreatedBy uint16 // User.ID
	UpdatedBy uint16 // User.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeleteInfo struct {
	DeletedBy *uint16 // User.ID
	DeletedAt time.Time
}
