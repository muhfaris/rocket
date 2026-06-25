package portservice

import (
	"context"

	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
)

//go:generate mockgen -destination=../../../../../shared/mock/service/service.go -package=mockservices -source=service.go
type BookSvc interface {
	GetBook(ctx context.Context, payload domain.GetBook) (domain.GetBook, error)
	BorrowBook(ctx context.Context, payload domain.BorrowBook) (domain.BorrowBook, error)
	ReturnBook(ctx context.Context, payload domain.ReturnBook) error
	HealthCheck(ctx context.Context, payload domain.HealthCheck) (domain.HealthCheck, error)
	ListBooks(ctx context.Context, payload domain.ListBooks) (domain.ListBooks, error)
	CreateBook(ctx context.Context, payload domain.CreateBook) (domain.CreateBook, error)
}
