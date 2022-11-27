package visitor

type VisitorInterface interface {
	Visit(Entry) any
}
