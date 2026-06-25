package service

import (
	"context"

	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
	portregistry "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/registry"
	portservice "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/service"
)

type BookSvc struct {
	reg portregistry.Repository
}

func NewBookSvc(reg portregistry.Repository) portservice.BookSvc {
	return &BookSvc{reg: reg}
}
func (s *BookSvc) HealthCheck(ctx context.Context, payload domain.HealthCheck) (domain.HealthCheck, error) {
	return domain.HealthCheck{}, nil
}
func (s *BookSvc) ListBooks(ctx context.Context, payload domain.ListBooks) (domain.ListBooks, error) {
	return domain.ListBooks{}, nil
}
func (s *BookSvc) CreateBook(ctx context.Context, payload domain.CreateBook) (domain.CreateBook, error) {
	return domain.CreateBook{}, nil
}
func (s *BookSvc) GetBook(ctx context.Context, payload domain.GetBook) (domain.GetBook, error) {
	return domain.GetBook{}, nil
}
func (s *BookSvc) BorrowBook(ctx context.Context, payload domain.BorrowBook) (domain.BorrowBook, error) {
	return domain.BorrowBook{}, nil
}
func (s *BookSvc) ReturnBook(ctx context.Context, payload domain.ReturnBook) error {
	return nil
}
