package bootstrap

import (
	portregistry "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/registry"
	portservice "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/service"
	"github.com/muhfaris/rocket/examples/samplepg/internal/core/service"
)

type AppService struct {
	reg portregistry.Repository
}

func InitializeService(reg portregistry.Repository) portregistry.Service {
	return &AppService{reg: reg}
}
func (a *AppService) GetReportSvc() portservice.ReportSvc {
	return service.NewReportSvc(a.reg)
}
