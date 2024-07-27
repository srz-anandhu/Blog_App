package repo

import "time"

// Blog Model
type Blog struct {
	ID        uint16
	Title     string
	Content   string
	Status    int    // 1 - Draft, 2 - Published, 3 - Deleted
	AuthorID  uint16 // Author.ID
	UpdatedBy uint16 // User.ID
	CreatedBy uint16 // User.ID
	DeletedBy uint16 // User.ID
	DeletedAt time.Time
	Model
}
