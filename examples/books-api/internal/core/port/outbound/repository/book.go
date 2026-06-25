package repository

import "context"

type BookRepository interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	FindByID(id string) error
	Find() error
}
