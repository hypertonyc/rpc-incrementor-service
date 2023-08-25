package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hypertonyc/rpc-incrementor-service/internal/repository"
)

func TestServiceGetNumber(t *testing.T) {
	repo := repository.NewInMemoryIncrementRepository()
	service := NewIncrementService(repo)

	// Default value in new repository should be zero
	number, err := service.GetNumber()
	require.NoError(t, err)
	assert.Equal(t, int32(0), number)
}

func TestServiceIncrement(t *testing.T) {
	repo := repository.NewInMemoryIncrementRepository()
	service := NewIncrementService(repo)

	err := service.IncrementNumber()
	require.NoError(t, err)
}

func TestServiceSetSettings(t *testing.T) {
	repo := repository.NewInMemoryIncrementRepository()
	service := NewIncrementService(repo)

	err := service.SetSettings(1, 10)
	require.NoError(t, err)
}

func TestServiceSetIncorrectIncrementStep(t *testing.T) {
	repo := repository.NewInMemoryIncrementRepository()
	service := NewIncrementService(repo)

	err := service.SetSettings(0, 10)
	require.Error(t, err)
}

func TestServiceSetIncorrectUpperLimit(t *testing.T) {
	repo := repository.NewInMemoryIncrementRepository()
	service := NewIncrementService(repo)

	err := service.SetSettings(1, 0)
	require.Error(t, err)
}

func TestServiceUpperLimitOverflow(t *testing.T) {
	repo := repository.NewInMemoryIncrementRepository()
	service := NewIncrementService(repo)

	// Default value in new repository should be zero
	number, err := service.GetNumber()
	require.NoError(t, err)
	assert.Equal(t, int32(0), number)

	// Set increment step to 3 and upper limit to 5
	err = service.SetSettings(3, 5)
	require.NoError(t, err)

	// Executing first increment
	err = service.IncrementNumber()
	require.NoError(t, err)

	// After first increment value should be 3
	number, err = service.GetNumber()
	require.NoError(t, err)
	assert.Equal(t, int32(3), number)

	// Executing second increment
	err = service.IncrementNumber()
	require.NoError(t, err)

	// After second increment value should be 1 (3 + 3 = 6; 6 - 5 = 1)
	number, err = service.GetNumber()
	require.NoError(t, err)
	assert.Equal(t, int32(1), number)
}
