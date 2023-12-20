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
	FindById(ctx context.Context, id string) (web.ProductGetResponse, error)
	Insert(ctx context.Context, product web.ProductNewRequest) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, product web.ProductUpdateRequest) error
	FindAll(ctx context.Context) ([]web.ProductListResponse, error)
	FindAllSortByTitle(ctx context.Context) ([]web.ProductListResponse, error)
	FindAllSortByRating(ctx context.Context) ([]web.ProductListResponse, error)
}

func NewProduct(db *sql.DB) productRepo {
	return &productImpl{DB: db}
}

func (p *productImpl) FindById(ctx context.Context, id string) (web.ProductGetResponse, error) {
	query := `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM Product WHERE id = ? LIMIT 1`
	rows, err := p.DB.QueryContext(ctx, query, id)
	product := web.ProductGetResponse{}
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

func (p *productImpl) Insert(ctx context.Context, product web.ProductNewRequest) error {
	query := `INSERT INTO Product(id, title, description, rating, image) VALUES (?, ?, ?, ?, ?)`
	_, err := p.DB.ExecContext(ctx, query,
		product.Id,
		product.Title,
		product.Description,
		product.Rating,
		product.Image,
	)
	return err
}

func (p *productImpl) Delete(ctx context.Context, id string) error {
	query := `UPDATE Product SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := p.DB.ExecContext(ctx, query, id)
	return err
}

func (p *productImpl) Update(ctx context.Context, product web.ProductUpdateRequest) error {
	query := `UPDATE Product SET title = ?, description = ?, rating = ?, image = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := p.DB.ExecContext(ctx, query,
		product.Title,
		product.Description,
		product.Rating,
		product.Image,
		product.Id,
	)
	return err
}

func (p *productImpl) FindAll(ctx context.Context) ([]web.ProductListResponse, error) {
	query := `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM Product`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return []web.ProductListResponse{}, err
	}
	var productList []web.ProductListResponse
	for rows.Next() {
		product := web.ProductListResponse{}
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

		productList = append(productList, product)
	}
	return productList, nil
}

func (p *productImpl) FindAllSortByTitle(ctx context.Context) ([]web.ProductListResponse, error) {
	query := `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM Product ORDER BY title ASC`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return []web.ProductListResponse{}, err
	}
	var productList []web.ProductListResponse
	for rows.Next() {
		product := web.ProductListResponse{}
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
		productList = append(productList, product)
	}
	return productList, nil
}

func (p *productImpl) FindAllSortByRating(ctx context.Context) ([]web.ProductListResponse, error) {
	query := `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at FROM Product ORDER BY rating ASC`
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return []web.ProductListResponse{}, err
	}
	var productList []web.ProductListResponse
	for rows.Next() {
		product := web.ProductListResponse{}
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
		productList = append(productList, product)
	}
	return productList, nil
}
