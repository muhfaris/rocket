package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
	"github.com/muhfaris/rocket/examples/books-api/shared/apierror"
)

type ListBooks struct{}

type ListBooksQueryQueryQuery struct {
	Status string `query:"status"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}

type ListBooksResponse struct {
	Items []ListBookItem `json:"items"`
	Total int            `json:"total"`
}

type ListBookItem struct {
	Author string `json:"author"`
	ID     string `json:"id"`
	Title  string `json:"title"`
}

func (req *ListBooks) In(c *fiber.Ctx) (domain.ListBooks, error) {

	var offset ListBooksQueryQueryQuery
	if err := c.QueryParser(&offset); err != nil {
		return domain.ListBooks{}, apierror.NewBadRequest("invalid request query", err)
	}

	return domain.ListBooks{

		Status: offset.Status,
		Limit:  offset.Limit,
		Offset: offset.Offset,
	}, nil
}

func (req *ListBooks) Out(c *fiber.Ctx, data domain.ListBooks) any {
	// TODO: map domain.ListBooks to ListBooksResponse
	return data
}
