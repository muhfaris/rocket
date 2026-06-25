package repository

import "context"

type ReportRepository interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	FindByID(id string) error
	Find() error
}
