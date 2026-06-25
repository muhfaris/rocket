package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/response"
) // POST /books handler
// @Summary Create a new book
// @Tags [Books]
// @Accept application/json
// @Param body body &amp;[object] true &#34;Request body&#34;
// @Success 201 {object}  &#34;Book created successfully&#34;
// @Router /books [POST]

func (h *Handler) CreateBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetBookSvc()
		)

		payload, err := new(presenter.CreateBook).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.CreateBook(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, new(presenter.CreateBook).Out(c, result))
	}
}
