package repository

import (
	"github.com/hypertonyc/rpc-incrementor-service/internal/domain"
)

type IncrementRepository interface {
	GetNumber() (*domain.Number, error)
	UpdateNumber(number *domain.Number) error
	GetSettings() (*domain.Settings, error)
	UpdateSettings(settings *domain.Settings) error
}
