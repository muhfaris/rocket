package portregistry

import portservice "github.com/muhfaris/rocket-examples/internal/core/port/inbound/service"

type Service interface {
	PartnerSvc() portservice.PartnerSvc
}
