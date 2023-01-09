package gostate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	timer := NewTimer()
	workState := NewWorkStateBehavior(timer)
	timer.ChangeState(workState)

	timer.Start()
	assert.Equal(t, 25*60, timer.Second)

	timer.Finish()
	assert.Equal(t, 0, timer.Second)
	assert.IsType(t, NewRestStateBehavior(timer), timer.StateBehavior)

	timer.Start()
	assert.Equal(t, 10*60, timer.Second)
}
