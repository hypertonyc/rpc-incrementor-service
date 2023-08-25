package repository

import (
	"testing"

	"github.com/hypertonyc/rpc-incrementor-service/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGetNumber(t *testing.T) {
	var err error
	number := new(domain.Number)

	repo := NewInMemoryIncrementRepository()
	err = repo.UpdateNumber(number)
	require.NoError(t, err)

	repoNumber, err := repo.GetNumber()
	require.NoError(t, err)

	assert.Equal(t, number, repoNumber)
}

func TestSetGetSettings(t *testing.T) {
	var err error
	settings := new(domain.Settings)

	repo := NewInMemoryIncrementRepository()
	err = repo.UpdateSettings(settings)
	require.NoError(t, err)

	repoSettings, err := repo.GetSettings()
	require.NoError(t, err)

	assert.Equal(t, settings, repoSettings)
}
