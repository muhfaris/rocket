package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/samplepg/internal/core/domain"
	"github.com/muhfaris/rocket/examples/samplepg/shared/apierror"
)

type GetReports struct{}

type GetReportsQueryQueryQuery struct {
	Status string `query:"status"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}

type ResponseGetReports []Report

type Report struct {
	ID          string `json:"id"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
}

func (req *GetReports) In(c *fiber.Ctx) (domain.GetReports, error) {

	var offset GetReportsQueryQueryQuery
	if err := c.QueryParser(&offset); err != nil {
		return domain.GetReports{}, apierror.NewBadRequest("invalid request query", err)
	}

	return domain.GetReports{

		Status: offset.Status,
		Limit:  offset.Limit,
		Offset: offset.Offset,
	}, nil
}

func (req *GetReports) Out(c *fiber.Ctx, data domain.GetReports) any {
	return ResponseGetReports{}
}
