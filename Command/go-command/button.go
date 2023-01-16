package gocommand

type Button struct {
	Command Command
}

func (b *Button) Click() {
	b.Command.Execute()
}
