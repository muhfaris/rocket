package group

import (
	"github.com/gofiber/fiber/v2"
	handlerv1 "github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/handler"
)

func V1(r *fiber.App, h *handlerv1.Handler) {
	// public
	// publicGroup := r.Group("/")
	// publicGroup.Get("/health", handlerv1.Health())

	routeGroup(r, h)
}

func routeGroup(r *fiber.App, h *handlerv1.Handler) {
	routeGroup := r.Group("/")
	routeGroup.Get("/reports", h.GetReports()).Name("GetReports")
	routeGroup.Post("/reports", h.CreateReport()).Name("CreateReport")
}
