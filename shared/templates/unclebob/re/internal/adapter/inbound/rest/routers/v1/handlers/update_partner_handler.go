package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-examples/internal/adapter/inbound/rest/routers/v1/response"
)

type UpdatePartner struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

func (h *Handler) UpdatePartnerHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.PartnerSvc()
				)
		*/

		var bodyRequest UpdatePartner
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.UpdatePartner(ctx , bodyRequest )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
