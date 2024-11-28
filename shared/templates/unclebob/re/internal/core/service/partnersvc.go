package service

import (
	"context"

	"github.com/muhfaris/rocket-examples/internal/core/domain"
	portservice "github.com/muhfaris/rocket-examples/internal/core/port/inbound/service"
)

type PartnerSvc struct{}

func NewPartnerSvc() portservice.PartnerSvc {
	return &PartnerSvc{}
}
func (s *PartnerSvc) GetPartners(ctx context.Context, bodyRequest map[string]any) error {
	return nil
}
func (s *PartnerSvc) GetDetailPartner(ctx context.Context, partner_id domain.DetailPartner) error {
	return nil
}
func (s *PartnerSvc) UpdatePartner(ctx context.Context, bodyRequest domain.UpdatePartner) error {
	return nil
}
