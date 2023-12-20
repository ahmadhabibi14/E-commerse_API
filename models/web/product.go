package web

import (
	"database/sql"
	"time"
)

type ProductGetRequest struct {
	Id string `validate:"required,min=35,max=36" json:"id"`
}

type ProductGetResponse struct {
	Id          string       `db:"id" json:"id"`
	Title       string       `db:"title" json:"title"`
	Description string       `db:"description" json:"description"`
	Rating      float64      `db:"rating" json:"rating"`
	Image       string       `db:"image" json:"image"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type ProductNewRequest struct {
	Id          string  `db:"id" json:"id"`
	Title       string  `db:"title" json:"title" validate:"required"`
	Description string  `db:"description" json:"description" validate:"required"`
	Rating      float64 `db:"rating" json:"rating"`
	Image       string  `db:"image" json:"image"`
}

type ProductDeleteRequest struct {
	Id string `validate:"required,min=35,max=36" json:"id"`
}

type ProductUpdateRequest struct {
	Id          string  `db:"id" json:"id" validate:"required,min=35,max=36"`
	Title       string  `db:"title" json:"title" validate:"required"`
	Description string  `db:"description" json:"description" validate:"required"`
	Rating      float64 `db:"rating" json:"rating"`
	Image       string  `db:"image" json:"image"`
}

type ProductListResponse struct {
	Id          string       `db:"id" json:"id"`
	Title       string       `db:"title" json:"title"`
	Description string       `db:"description" json:"description"`
	Rating      float64      `db:"rating" json:"rating"`
	Image       string       `db:"image" json:"image"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"deleted_at"`
}
