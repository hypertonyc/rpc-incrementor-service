package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hypertonyc/rpc-incrementor-service/internal/domain"
)

type pgIncrementRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewPostgresIncrementRepository(db *sql.DB, ctx context.Context) IncrementRepository {
	return &pgIncrementRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *pgIncrementRepository) GetNumber() (*domain.Number, error) {
	var value int32
	query := "SELECT value FROM numbers LIMIT 1"
	err := r.db.QueryRowContext(r.ctx, query).Scan(&value)
	if err != nil {
		return nil, fmt.Errorf("error querying number: %w", err)
	}

	var number domain.Number
	number.SetValue(value)

	return &number, nil
}

func (r *pgIncrementRepository) UpdateNumber(number *domain.Number) error {
	query := "UPDATE numbers SET value = $1"
	_, err := r.db.ExecContext(r.ctx, query, number.GetValue())
	if err != nil {
		return fmt.Errorf("error updating number: %w", err)
	}

	return nil
}

func (r *pgIncrementRepository) GetSettings() (*domain.Settings, error) {
	var incrementStep, upperLimit int32
	query := "SELECT increment_step, upper_limit FROM settings LIMIT 1"
	err := r.db.QueryRowContext(r.ctx, query).Scan(&incrementStep, &upperLimit)
	if err != nil {
		return nil, fmt.Errorf("error querying settings: %w", err)
	}

	var settings domain.Settings
	settings.SetUpperLimit(upperLimit)
	settings.SetIncrementStep(incrementStep)

	return &settings, nil
}

func (r *pgIncrementRepository) UpdateSettings(settings *domain.Settings) error {
	query := "UPDATE settings SET increment_step = $1, upper_limit = $2"
	_, err := r.db.ExecContext(r.ctx, query, settings.GetIncrementStep(), settings.GetUpperLimit())
	if err != nil {
		return fmt.Errorf("error updating settings: %w", err)
	}

	return nil
}
