package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Price       int
	Category_id uuid.UUID
	Created_at  string
	Upated_at   string
}
