package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
	"github.com/muhfaris/rocket/examples/books-api/shared/apierror"
)

type BorrowBook struct{}

type BorrowBookParams struct {
	Bookid string `params:"bookId"`
}
type BorrowBookRequest struct {
	MemberID string `json:"member_id"`
}

type BorrowBookResponse struct {
	ID      string `json:"id"`
	DueDate string `json:"due_date"`
}

func (req *BorrowBook) In(c *fiber.Ctx) (domain.BorrowBook, error) {

	var borrowBookParams BorrowBookParams
	if err := c.ParamsParser(&borrowBookParams); err != nil {
		return domain.BorrowBook{}, apierror.NewBadRequest("invalid request params", err)
	}

	var bodyRequest BorrowBookRequest
	if err := c.BodyParser(&bodyRequest); err != nil {
		return domain.BorrowBook{}, apierror.NewBadRequest("invalid request body", err)
	}

	return domain.BorrowBook{

		Bookid: borrowBookParams.Bookid,

		MemberID: bodyRequest.MemberID,
	}, nil
}

func (req *BorrowBook) Out(c *fiber.Ctx, data domain.BorrowBook) any {
	return BorrowBookResponse{
		ID: data.ID,
	}
}
