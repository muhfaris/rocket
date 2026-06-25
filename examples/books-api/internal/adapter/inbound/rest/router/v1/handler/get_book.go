package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/response"
) // GET /books/{bookId} handler
// @Summary Get a book by ID
// @Tags [Books]
// @Param bookId path &amp;[string] true &#34;&#34;
// @Success 200 {object}  &#34;A book&#34;
// @Router /books/{bookId} [GET]

func (h *Handler) GetBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetBookSvc()
		)

		payload, err := new(presenter.GetBook).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.GetBook(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, new(presenter.GetBook).Out(c, result))
	}
}
