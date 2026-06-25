package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
	"github.com/muhfaris/rocket/examples/books-api/shared/apierror"
)

type CreateBook struct{}

type CreateBookRequest struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookResponse struct {
	ID string `json:"id"`
}

func (req *CreateBook) In(c *fiber.Ctx) (domain.CreateBook, error) {

	var bodyRequest CreateBookRequest
	if err := c.BodyParser(&bodyRequest); err != nil {
		return domain.CreateBook{}, apierror.NewBadRequest("invalid request body", err)
	}

	return domain.CreateBook{

		Isbn:   bodyRequest.Isbn,
		Title:  bodyRequest.Title,
		Author: bodyRequest.Author,
	}, nil
}

func (req *CreateBook) Out(c *fiber.Ctx, data domain.CreateBook) any {
	return CreateBookResponse{
		ID: data.ID,
	}
}
