package model

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/google/uuid"
	"time"
)

type Article struct {
	tableName struct{}      `pg:"article"`
	Id        uuid.UUID     `json:"id"  pg:",pk,type:uuid"`
	Title     string        `json:"title" validate:"required"`
	Content   string        `json:"content" validate:"required"`
	AuthorID  uuid.UUID     `json:"author_id" pg:"type:uuid" validate:"authorMustExist"`
	Author    *model.Author `json:"author" pg:"rel:has-one"`
	CreatedAt time.Time     `json:"created_at" `
	UpdatedAt time.Time     `json:"updated_at" `
	DeletedAt time.Time     `json:"deleted_at"  pg:",soft_delete"`
}

type FetchParam struct {
	Title    *string `query:"title"`
	AuthorID *string `query:"author_id"`
	From     *string `query:"from"`
	Until    *string `query:"until"`
	Page     *int    `query:"page"`
	Size     *int    `query:"size"`
	Paginate *bool   `query:"paginate"`
}

type FetchResponse struct {
	Id        uuid.UUID            `json:"id"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Author    *model.FetchResponse `json:"author"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
