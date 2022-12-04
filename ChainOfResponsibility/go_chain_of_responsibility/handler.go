package cor

type RequestHandlerInterface interface {
	Handle(ServerRequestInterface) ResponseInterface
}

type DoubleRequestHandler struct {
}

func (d *DoubleRequestHandler) Handle(s ServerRequestInterface) ResponseInterface {
	content := s.GetBody()
	r := NewResponse(content + content)
	return r
}
