package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettingsGetterSetter(t *testing.T) {
	var err error
	settings := Settings{}

	err = settings.SetUpperLimit(10)
	require.NoError(t, err)
	assert.Equal(t, int32(10), settings.GetUpperLimit())

	err = settings.SetIncrementStep(1)
	require.NoError(t, err)
	assert.Equal(t, int32(1), settings.GetIncrementStep())
}

func TestIncrementStepZeroValue(t *testing.T) {
	settings := Settings{}
	err := settings.SetIncrementStep(0)

	require.ErrorIs(t, err, ErrStepZeroOrLess)
}

func TestIncrementStepNegativeValue(t *testing.T) {
	settings := Settings{}
	err := settings.SetIncrementStep(-10)

	require.ErrorIs(t, err, ErrStepZeroOrLess)
}

func TestUpperLimitZeroValue(t *testing.T) {
	settings := Settings{}
	err := settings.SetUpperLimit(0)

	require.ErrorIs(t, err, ErrLimitZeroOrLess)
}

func TestUpperLimitNegativeValue(t *testing.T) {
	settings := Settings{}
	err := settings.SetUpperLimit(-10)

	require.ErrorIs(t, err, ErrLimitZeroOrLess)
}

func TestStepGreaterThanLimit(t *testing.T) {
	settings := Settings{}

	settings.SetUpperLimit(10)
	err := settings.SetIncrementStep(20)

	require.ErrorIs(t, err, ErrStepGreaterThanLimit)
}

func TestLimitLessThanStep(t *testing.T) {
	settings := Settings{}

	settings.SetUpperLimit(10)
	settings.SetIncrementStep(5)

	err := settings.SetUpperLimit(3)

	require.ErrorIs(t, err, ErrLimitLessThanStep)
}
