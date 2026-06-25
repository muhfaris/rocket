package portservice

import (
	"context"

	"github.com/muhfaris/rocket/examples/samplepg/internal/core/domain"
)

//go:generate mockgen -destination=../../../../../shared/mock/service/service.go -package=mockservices -source=service.go
type ReportSvc interface {
	GetReports(ctx context.Context, payload domain.GetReports) (domain.GetReports, error)
	CreateReport(ctx context.Context, payload domain.CreateReport) (domain.CreateReport, error)
}
