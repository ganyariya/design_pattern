package gocommand

type Canvas struct {
	drawnArea string
}

func (c *Canvas) Draw(anything string) {
	c.drawnArea += anything
}

func (c *Canvas) Reset() {
	c.drawnArea = ""
}

func (c *Canvas) GetDrawnArea() string {
	return c.drawnArea
}
