package cor

type RequestHandlerInterface interface {
	handle(ServerRequestInterface) ResponseInterface
}

type DoubleRequestHandler struct {
}

func (d *DoubleRequestHandler) handle(s ServerRequestInterface) ResponseInterface {
	content := s.GetBody()
	r := NewResponse(content + content)
	return r
}
