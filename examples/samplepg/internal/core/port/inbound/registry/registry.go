package portregistry

import (
	portservice "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/service"
	"github.com/muhfaris/rocket/examples/samplepg/internal/core/port/outbound/repository"
)

//go:generate mockgen -destination=../../../../../shared/mock/registry/registry.go -package=mockregistry -source=registry.go
type Service interface {
	GetReportSvc() portservice.ReportSvc
}

type Repository interface {
	GetReportRepository() repository.ReportRepository
}
