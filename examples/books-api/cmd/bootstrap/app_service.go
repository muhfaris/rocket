package bootstrap

import (
	portregistry "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/registry"
	portservice "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/service"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/service"
)

type AppService struct {
	reg portregistry.Repository
}

func InitializeService(reg portregistry.Repository) portregistry.Service {
	return &AppService{reg: reg}
}
func (a *AppService) GetBookSvc() portservice.BookSvc {
	return service.NewBookSvc(a.reg)
}
