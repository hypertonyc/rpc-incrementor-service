package domain

import "errors"

var ErrStepZeroOrLess = errors.New("increment step must be greater than zero")
var ErrStepGreaterThanLimit = errors.New("increment step must be less than upper limit")
var ErrLimitZeroOrLess = errors.New("upper limit must be greater than zero")
var ErrLimitLessThanStep = errors.New("upper limit must be greater than increment step")

type Settings struct {
	upperLimit    int32
	incrementStep int32
}

func (s *Settings) GetUpperLimit() int32 {
	return s.upperLimit
}

func (s *Settings) GetIncrementStep() int32 {
	return s.incrementStep
}

func (s *Settings) SetUpperLimit(limit int32) error {
	if limit <= 0 {
		return ErrLimitZeroOrLess
	}

	if limit <= s.incrementStep {
		return ErrLimitLessThanStep
	}

	s.upperLimit = limit
	return nil
}

func (s *Settings) SetIncrementStep(step int32) error {
	if step <= 0 {
		return ErrStepZeroOrLess
	}

	if step >= s.upperLimit {
		return ErrStepGreaterThanLimit
	}

	s.incrementStep = step
	return nil
}
