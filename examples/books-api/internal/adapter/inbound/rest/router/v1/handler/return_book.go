package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/response"
) // PATCH /books/{bookId}/return handler
// @Summary Return a borrowed book
// @Tags [Books]
// @Param bookId path &amp;[string] true &#34;&#34;
// @Success 200 {object}  &#34;Book returned successfully&#34;
// @Router /books/{bookId}/return [PATCH]

func (h *Handler) ReturnBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetBookSvc()
		)

		payload, err := new(presenter.ReturnBook).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		err = svc.ReturnBook(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, "Hello")
	}
}
