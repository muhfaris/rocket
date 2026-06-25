package group

import (
	"github.com/gofiber/fiber/v2"
	handlerv1 "github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/handler"
)

func V1(r *fiber.App, h *handlerv1.Handler) {
	// public
	// publicGroup := r.Group("/")
	// publicGroup.Get("/health", handlerv1.Health())

	routeGroup(r, h)

	bookGroup(r, h)

	borrowGroup(r, h)
}

func routeGroup(r *fiber.App, h *handlerv1.Handler) {
	routeGroup := r.Group("/")
	routeGroup.Get("/health", h.HealthCheck()).Name("HealthCheck")
}

func bookGroup(r *fiber.App, h *handlerv1.Handler) {
	bookGroup := r.Group("/api/v1")
	bookGroup.Get("/books", h.ListBooks()).Name("ListBooks")
	bookGroup.Post("/books", h.CreateBook()).Name("CreateBook")
	bookGroup.Get("/books/:bookId", h.GetBook()).Name("GetBook")
}

func borrowGroup(r *fiber.App, h *handlerv1.Handler) {
	borrowGroup := r.Group("/api/v1")
	borrowGroup.Post("/books/:bookId/borrow", h.BorrowBook()).Name("BorrowBook")
	borrowGroup.Patch("/books/:bookId/return", h.ReturnBook()).Name("ReturnBook")
}
