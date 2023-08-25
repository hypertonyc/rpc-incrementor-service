package repository

import (
	"github.com/hypertonyc/rpc-incrementor-service/internal/domain"
)

type inMemoryIncrementRepository struct {
	Number        int32
	IncrementStep int32
	UpperLimit    int32
}

func NewInMemoryIncrementRepository() IncrementRepository {
	return &inMemoryIncrementRepository{
		Number:        0,
		IncrementStep: 1,
		UpperLimit:    10,
	}
}

func (ir *inMemoryIncrementRepository) GetNumber() (*domain.Number, error) {
	var number domain.Number
	number.SetValue(ir.Number)

	return &number, nil
}

func (ir *inMemoryIncrementRepository) UpdateNumber(number *domain.Number) error {
	ir.Number = number.GetValue()

	return nil
}

func (ir *inMemoryIncrementRepository) GetSettings() (*domain.Settings, error) {
	var settings domain.Settings

	settings.SetUpperLimit(ir.UpperLimit)
	settings.SetIncrementStep(ir.IncrementStep)

	return &settings, nil
}

func (ir *inMemoryIncrementRepository) UpdateSettings(settings *domain.Settings) error {
	ir.UpperLimit = settings.GetUpperLimit()
	ir.IncrementStep = settings.GetIncrementStep()

	return nil
}
