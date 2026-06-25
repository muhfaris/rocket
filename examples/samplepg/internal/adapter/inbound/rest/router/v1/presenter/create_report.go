package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/samplepg/internal/core/domain"
	"github.com/muhfaris/rocket/examples/samplepg/shared/apierror"
)

type CreateReport struct{}

type CreateReportRequest struct {
	Description string `json:"description"`
	Location    string `json:"location"`
	Title       string `json:"title"`
}

type Data struct {
	ID string `json:"id"`
}

func (req *CreateReport) In(c *fiber.Ctx) (domain.CreateReport, error) {

	var bodyRequest CreateReportRequest
	if err := c.BodyParser(&bodyRequest); err != nil {
		return domain.CreateReport{}, apierror.NewBadRequest("invalid request body", err)
	}

	return domain.CreateReport{

		Description: bodyRequest.Description,
		Location:    bodyRequest.Location,
		Title:       bodyRequest.Title,
	}, nil
}

func (req *CreateReport) Out(c *fiber.Ctx, data domain.CreateReport) any {
	return Data{
		ID: data.ID,
	}
}
