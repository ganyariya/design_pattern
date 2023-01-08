package gomediator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDispatcher(t *testing.T) {
	dispatcher := NewEventDispatcher()

	listenerForm := NewEventListenerForm(dispatcher, "submit")
	listenerButton := NewEventListenerButton(dispatcher, "*")

	listenerButton.Click("submit")
	assert.Equal(t, "button submitted", listenerForm.NotifiedMessage)
}
