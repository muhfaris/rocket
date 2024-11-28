package portservice

import (
	"context"

	"github.com/muhfaris/rocket-examples/internal/core/domain"
)

type PartnerSvc interface {
	GetPartners(ctx context.Context, bodyRequest map[string]any) error
	GetDetailPartner(ctx context.Context, partner_id domain.DetailPartner) error
	UpdatePartner(ctx context.Context, bodyRequest domain.UpdatePartner) error
}
