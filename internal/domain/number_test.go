package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetterSetter(t *testing.T) {
	number := Number{}
	number.SetValue(10)
	assert.Equal(t, int32(10), number.GetValue())
}

func TestIncrement(t *testing.T) {
	number := Number{}
	number.SetValue(10)
	number.Increment(1, 100)
	assert.Equal(t, int32(11), number.GetValue())
}

func TestIncrementUpperLimit(t *testing.T) {
	number := Number{}
	number.SetValue(4)
	number.Increment(1, 5)
	assert.Equal(t, int32(0), number.GetValue())
}

func TestIncrementUpperLimitOverflow(t *testing.T) {
	number := Number{}
	number.SetValue(4)
	number.Increment(3, 5)
	assert.Equal(t, int32(2), number.GetValue())
}
