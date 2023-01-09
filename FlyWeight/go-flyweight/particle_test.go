package goflyweight

import (
	"testing"

	"github.com/ganyariya/design_pattern/FlyWeight/go-flyweight/particle"
	"github.com/stretchr/testify/assert"
)

func TestParticle(t *testing.T) {
	initX, initY := 10, 20
	color := "green"
	image := "tama.png"

	p1 := particle.NewMovingParticle(initX, initY, color, image)
	p1.Move(10, 20)

	assert.Equal(t, "20 40 green tama.png", p1.Draw())
}
