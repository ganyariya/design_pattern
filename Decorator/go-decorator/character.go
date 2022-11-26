package decorator

type CharacterInterface interface {
	Draw() string
}

type Character struct {
	Name string
}

func (c *Character) Draw() string {
	return c.Name
}

func NewCharacter(name string) *Character {
	return &Character{name}
}
