package types

type EntityTime struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"-"`
}

type BaseEntity struct {
	ID        ID    `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"-"`
}
