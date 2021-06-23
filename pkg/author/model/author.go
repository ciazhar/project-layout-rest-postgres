package model

import (
	"github.com/google/uuid"
	"time"
)

type Author struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type FetchParam struct {
	Name     *string `query:"name"`
	Page     *int    `query:"page"`
	Size     *int    `query:"size"`
	Paginate *bool   `query:"paginate"`
}

type FetchResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
