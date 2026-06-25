package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/response"
) // POST /reports handler
// @Summary Submit a new report
// @Tags [Reports]
// @Accept application/json
// @Param body body &amp;[object] true &#34;Request body&#34;
// @Success 201 {object}  &#34;Report created successfully&#34;
// @Router /reports [POST]

func (h *Handler) CreateReport() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetReportSvc()
		)

		payload, err := new(presenter.CreateReport).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.CreateReport(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, result)
	}
}
