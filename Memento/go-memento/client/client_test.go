package client

import (
	"testing"

	"github.com/ganyariya/design_pattern/Memento/go-memento/game"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	originator := game.NewOriginator(0)
	assert.Equal(t, 0, originator.GetCounter())

	memento := originator.CreateMemento()
	for i := 0; i < 10; i++ {
		originator.Play()
	}

	assert.Equal(t, 10, originator.GetCounter())

	originator.RestoreMemento(memento)

	assert.Equal(t, 0, originator.GetCounter())
}
