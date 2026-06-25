package pgsqlrepository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/port/outbound/repository"
)

type BookRepository struct {
	conn *pgxpool.Pool
}

func NewBookRepository(conn repository.PSQLRepository) repository.BookRepository {
	return &BookRepository{conn: conn.GetConnection()}
}

func (r *BookRepository) Create(ctx context.Context) error {
	return nil
}

func (r *BookRepository) Update(ctx context.Context) error {
	return nil
}

func (r *BookRepository) FindByID(id string) error {
	return nil
}

func (r *BookRepository) Find() error {
	return nil
}
