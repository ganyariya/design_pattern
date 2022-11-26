package decorator

import "fmt"

type Sleep struct {
	c CharacterInterface
}

func (s *Sleep) Draw() string {
	return fmt.Sprintf("%s: sleep", s.c.Draw())
}
