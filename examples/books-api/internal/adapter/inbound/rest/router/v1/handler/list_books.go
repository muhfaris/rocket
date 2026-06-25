package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/response"
) // GET /books handler
// @Summary List all books
// @Tags [Books]
// @Param status query &amp;[string] false &#34;&#34;
// @Param limit query &amp;[integer] false &#34;&#34;
// @Param offset query &amp;[integer] false &#34;&#34;
// @Success 200 {object}  &#34;A list of books&#34;
// @Router /books [GET]

func (h *Handler) ListBooks() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetBookSvc()
		)

		payload, err := new(presenter.ListBooks).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.ListBooks(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, new(presenter.ListBooks).Out(c, result))
	}
}
