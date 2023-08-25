package domain

type Number struct {
	value int32
}

func (n *Number) GetValue() int32 {
	return n.value
}

func (n *Number) SetValue(value int32) {
	n.value = value
}

func (n *Number) Increment(incrementStep int32, upperLimit int32) {
	if n.value >= (upperLimit - incrementStep) {
		// calculating this way prevents int32 overflow
		restVal := upperLimit - incrementStep
		n.value -= restVal
	} else {
		n.value += incrementStep
	}
}
