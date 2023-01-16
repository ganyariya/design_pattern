package gocommand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	button := &Button{}
	canvas := &Canvas{}

	drawCommand := NewDrawCommand(canvas, "hello")
	button.Command = drawCommand

	button.Click()
	button.Click()

	assert.Equal(t, "hellohello", canvas.GetDrawnArea())

	resetCommand := NewResetCommand(canvas)
	button.Command = resetCommand

	button.Click()
	assert.Equal(t, "", canvas.GetDrawnArea())
}
