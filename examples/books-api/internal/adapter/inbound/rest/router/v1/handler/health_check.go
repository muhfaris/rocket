package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router/v1/response"
) // GET /health handler
// @Summary Health check endpoint
// @Tags [Books]
// @Success 200 {object}  &#34;OK&#34;
// @Router /health [GET]

func (h *Handler) HealthCheck() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetBookSvc()
		)

		payload, err := new(presenter.HealthCheck).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.HealthCheck(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, new(presenter.HealthCheck).Out(c, result))
	}
}
