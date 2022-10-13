package go_builder

type RedBuilder struct {
	rect Rect
}

func (r *RedBuilder) New(x, y int) {
	r.rect = Rect{}
	r.rect.x = x
	r.rect.y = y
}

func (r *RedBuilder) SetColor(c string) {
	r.rect.color = c
}
func (r *RedBuilder) SetFill(f bool) {
	r.rect.fill = f
}
func (r *RedBuilder) SetSize(w, h int) {
	r.rect.w = w
	r.rect.h = h
}

func (r *RedBuilder) Get() Rect {
	return r.rect
}
