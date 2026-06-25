package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket/examples/books-api/internal/core/domain"
)

type HealthCheck struct{}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func (req *HealthCheck) In(c *fiber.Ctx) (domain.HealthCheck, error) {

	return domain.HealthCheck{}, nil
}

func (req *HealthCheck) Out(c *fiber.Ctx, data domain.HealthCheck) any {
	// TODO: map domain.HealthCheck to HealthCheckResponse
	return data
}
