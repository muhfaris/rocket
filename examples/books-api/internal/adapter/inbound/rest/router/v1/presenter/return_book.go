package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
	"github.com/muhfaris/rocket/examples/books-api/shared/apierror"
)

type ReturnBook struct{}

type ReturnBookParams struct {
	Bookid string `params:"bookId"`
}

func (req *ReturnBook) In(c *fiber.Ctx) (domain.ReturnBook, error) {

	var returnBookParams ReturnBookParams
	if err := c.ParamsParser(&returnBookParams); err != nil {
		return domain.ReturnBook{}, apierror.NewBadRequest("invalid request params", err)
	}

	return domain.ReturnBook{

		Bookid: returnBookParams.Bookid,
	}, nil
}
