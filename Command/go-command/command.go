package gocommand

type Command interface {
	Execute()
}

type DrawCommand struct {
	canvas *Canvas
	txt    string
}

func NewDrawCommand(canvas *Canvas, txt string) *DrawCommand {
	return &DrawCommand{canvas: canvas, txt: txt}
}
func (dc *DrawCommand) Execute() {
	dc.canvas.Draw(dc.txt)
}

type ResetCommand struct {
	canvas *Canvas
}

func NewResetCommand(canvas *Canvas) *ResetCommand {
	return &ResetCommand{canvas: canvas}
}
func (rc *ResetCommand) Execute() {
	rc.canvas.Reset()
}
