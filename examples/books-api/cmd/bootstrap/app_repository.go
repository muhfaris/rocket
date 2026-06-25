package bootstrap

import (
	portregistry "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/registry"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/port/outbound/repository"
)

func InitializeRepository() portregistry.Repository {
	return &AppRepository{}
}

// AppRepository struct would need to be updated to include new repository types
type AppRepository struct {
}

func (a *AppRepository) GetBookRepository() repository.BookRepository {
	return nil
}
