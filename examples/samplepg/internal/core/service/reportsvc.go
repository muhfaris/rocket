package service

import (
	"context"

	"github.com/muhfaris/rocket/examples/samplepg/internal/core/domain"
	portregistry "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/registry"
	portservice "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/service"
)

type ReportSvc struct {
	reg portregistry.Repository
}

func NewReportSvc(reg portregistry.Repository) portservice.ReportSvc {
	return &ReportSvc{reg: reg}
}
func (s *ReportSvc) GetReports(ctx context.Context, payload domain.GetReports) (domain.GetReports, error) {
	return domain.GetReports{}, nil
}
func (s *ReportSvc) CreateReport(ctx context.Context, payload domain.CreateReport) (domain.CreateReport, error) {
	return domain.CreateReport{}, nil
}
