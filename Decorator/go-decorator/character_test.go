package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	var charaInterface CharacterInterface = nil

	charaInterface = NewCharacter("ganyariya")
	charaInterface = &Sleep{c: &Poison{c: charaInterface}}

	assert.Equal(t, "ganyariya: poison: sleep", charaInterface.Draw())
}
