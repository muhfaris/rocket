package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-examples/internal/adapter/inbound/rest/routers/v1/response"
)

type DetailPartner struct {
	PartnerID string `params:"partner_id"`
}

func (h *Handler) GetDetailPartner() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.PartnerSvc()
				)
		*/

		var partner_id DetailPartner
		if err := c.ParamsParser(&partner_id); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.GetDetailPartner(ctx , partner_id )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
