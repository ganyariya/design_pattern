package go_builder

type Director struct {
	builder Builder
}

func (d *Director) Construct() Rect {
	d.builder.New(10, 20)
	d.builder.SetColor("Yellow")
	return d.builder.Get()
}
