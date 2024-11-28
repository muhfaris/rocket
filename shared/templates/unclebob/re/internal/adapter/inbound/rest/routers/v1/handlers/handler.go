package handlersv1

import portregistry "github.com/muhfaris/rocket-examples/internal/core/port/inbound/registry"

type Handler struct {
	services portregistry.Service
}

func New() *Handler {
	return &Handler{}
}
