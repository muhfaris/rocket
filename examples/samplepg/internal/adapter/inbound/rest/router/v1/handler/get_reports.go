package handlerv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/presenter"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router/v1/response"
) // GET /reports handler
// @Summary Get all reports
// @Tags [Reports]
// @Param status query &amp;[string] false &#34;&#34;
// @Param limit query &amp;[integer] false &#34;&#34;
// @Param offset query &amp;[integer] false &#34;&#34;
// @Success 200 {object} #/components/schemas/ResponseGetReports &#34;OK&#34;
// @Router /reports [GET]

func (h *Handler) GetReports() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.UserContext()
			svc = h.services.GetReportSvc()
		)

		payload, err := new(presenter.GetReports).In(c)
		if err != nil {
			return err
		}

		// Transform request into domain model
		result, err := svc.GetReports(ctx, payload)
		if err != nil {
			return err
		}

		return response.Success(c, result)
	}
}
