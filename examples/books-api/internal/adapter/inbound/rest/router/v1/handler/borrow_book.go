package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/response"
) // POST /books/{bookId}/borrow handler
// @Summary Borrow a book
// @Tags [Books]
// @Param bookId path &amp;[string] true &#34;&#34;
// @Accept application/json
// @Param body body &amp;[object] true &#34;Request body&#34;
// @Success 201 {object}  &#34;Book borrowed successfully&#34;
// @Router /books/{bookId}/borrow [POST]

func (h *Handler) BorrowBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetBookSvc()
		)

		payload, err := new(presenter.BorrowBook).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.BorrowBook(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, new(presenter.BorrowBook).Out(c, result))
	}
}
