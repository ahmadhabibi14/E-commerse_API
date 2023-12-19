package data

import (
	"context"
	"database/sql"
	"e-commerse_api/models/web"
	"errors"
)

type productImpl struct {
	DB *sql.DB
}

type productRepo interface {
	FindById(ctx context.Context, id string) (web.ProductResponse, error)
}

func NewProduct(db *sql.DB) productRepo {
	return &productImpl{DB: db}
}

func (p *productImpl) FindById(ctx context.Context, id string) (web.ProductResponse, error) {
	query := `SELECT * FROM Product WHERE id = ? LIMIT 1`
	rows, err := p.DB.QueryContext(ctx, query, id)
	product := web.ProductResponse{}
	if err != nil {
		return product, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.Rating,
			&product.Image,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
	} else {
		return product, errors.New(`Product not found`)
	}

	return product, nil
}
