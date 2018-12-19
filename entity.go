package types

import "time"

type BaseEntity struct {
	ID        ID        `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
