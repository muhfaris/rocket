package pgsqlrepository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/muhfaris/rocket/examples/samplepg/internal/core/port/outbound/repository"
)

type ReportRepository struct {
	conn *pgxpool.Pool
}

func NewReportRepository(conn repository.PSQLRepository) repository.ReportRepository {
	return &ReportRepository{conn: conn.GetConnection()}
}

func (r *ReportRepository) Create(ctx context.Context) error {
	return nil
}

func (r *ReportRepository) Update(ctx context.Context) error {
	return nil
}

func (r *ReportRepository) FindByID(id string) error {
	return nil
}

func (r *ReportRepository) Find() error {
	return nil
}
