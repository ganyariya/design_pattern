package go_builder

type Builder interface {
	New(int, int)
	SetColor(string)
	SetFill(bool)
	SetSize(int, int)
	Get() Rect
}
