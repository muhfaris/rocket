package portregistry

import (
	portservice "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/service"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/port/outbound/repository"
)

//go:generate mockgen -destination=../../../../../shared/mock/registry/registry.go -package=mockregistry -source=registry.go
type Service interface {
	GetBookSvc() portservice.BookSvc
}

type Repository interface {
	CacheRepository() repository.CacheRepository
	GetBookRepository() repository.BookRepository
}
