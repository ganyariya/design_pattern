package decorator

import "fmt"

type Poison struct {
	c CharacterInterface
}

func (p *Poison) Draw() string {
	return fmt.Sprintf("%s: poison", p.c.Draw())
}
