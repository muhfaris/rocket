package handlerv1

import portregistry "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/registry"

type Handler struct {
	services portregistry.Service
}

func New(svcs portregistry.Service) *Handler {
	return &Handler{services: svcs}
}
