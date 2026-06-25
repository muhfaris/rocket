package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
	"github.com/muhfaris/rocket/examples/books-api/shared/apierror"
)

type GetBook struct{}

type GetBookParams struct {
	Bookid string `params:"bookId"`
}

type GetBookResponse struct {
	Isbn   string `json:"isbn"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ID     string `json:"id"`
}

func (req *GetBook) In(c *fiber.Ctx) (domain.GetBook, error) {

	var bookId GetBookParams
	if err := c.ParamsParser(&bookId); err != nil {
		return domain.GetBook{}, apierror.NewBadRequest("invalid request params", err)
	}

	return domain.GetBook{

		Bookid: bookId.Bookid,
	}, nil
}

func (req *GetBook) Out(c *fiber.Ctx, data domain.GetBook) any {
	// TODO: map domain.GetBook to GetBookResponse
	return data
}
